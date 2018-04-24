package cargo

import (
	"fmt"
)

const (
	optRateLimit       = "rate_limit"
	optWorkerBatchSize = "worker_batch_size"
	optCollectStats    = "collect_stats"
)

var (
	errUnknownOpt      = "pool: unknown opt: `%s`"
	errOptTypeMismatch = "pool: option type mismatch: given: `%T`, expected: `%T`"
)

type poolCfg struct {
	rateLimit    *int
	batchSize    *int
	collectStats bool
	handler      ResultHandler
}

func (c *poolCfg) Set(k string, v interface{}) (err error) {
	switch k {
	case optRateLimit:
		rateLimit, ok := v.(int)
		if !ok {
			err = fmt.Errorf(errOptTypeMismatch, v, *c.batchSize)
		}
		c.rateLimit = &rateLimit

	case optWorkerBatchSize:
		batchSize, ok := v.(int)
		if !ok {
			err = fmt.Errorf(errOptTypeMismatch, v, *c.batchSize)
		}
		c.batchSize = &batchSize
	case optCollectStats:
		c.collectStats = true

	default:
		err = fmt.Errorf(errUnknownOpt, k)
	}

	return
}

type Options interface {
	Set(string, interface{}) error
}

type Option func(Options) error

func RateLimit(max int) Option {
	return func(opts Options) error {
		return opts.Set(optRateLimit, max)
	}
}

func WorkerBatchSize(max int) Option {
	return func(opts Options) error {
		return opts.Set(optWorkerBatchSize, max)
	}
}

func WithStats() Option {
	return func(opts Options) error {
		return opts.Set(optCollectStats, true)
	}
}
