package cargo

import (
	"fmt"
	"sync"
	"time"
)

var (
	errPoolSizeNegative = fmt.Errorf("pool: pool size cannot be negative")
)

type Pool interface {
	Close()
	Execute(Job)
}

type PoolWStats interface {
	Stats()
}

type Limiter interface {
	Allow() bool
	Stop()
}

type rateLimit struct {
	limiter Limiter
	p       Pool
}

func (p *rateLimit) Close() {
	p.p.Close()
}

func (p *rateLimit) Execute(j Job) {
	for {
		if p.limiter.Allow() {
			p.p.Execute(j)
			return
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func newRateLimit(limit int32, p Pool) Pool {
	return &rateLimit{p: p, limiter: newLimiter(limit)}
}

type pool struct {
	workers chan Worker
	handler ResultHandler
	wg      *sync.WaitGroup
}

func (p *pool) Close() {
	p.wg.Wait()
}

func (p *pool) Execute(job Job) {
	worker := <-p.workers
	p.wg.Add(1)

	go func(p *pool, worker Worker) {
		defer func() {
			p.wg.Done()
			p.workers <- worker
		}()

		result := worker.Process(job)

		p.handler.Handle(result)
	}(p, worker)
}

func NewPool(
	size int,
	handler ResultHandler,
	newWorker func() Worker,
	opts ...Option,
) (_ Pool, err error) {
	if size < 0 {
		return nil, errPoolSizeNegative
	}

	cfg := &poolCfg{}

	for _, opt := range opts {
		if err = opt(cfg); err != nil {
			return
		}
	}

	p := pool{
		wg:      &sync.WaitGroup{},
		workers: make(chan Worker, size),
		handler: handler,
	}

	for i := 0; i < size; i++ {
		p.workers <- newWorker()
	}

	if cfg.rateLimit != nil {
		return newRateLimit(int32(*cfg.rateLimit), &p), nil
	}

	return &p, nil
}
