package cargo_test

import (
	"log"
	"time"

	"testing"
	"math/rand"

	"github.com/bmartynov/cargo"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

type job struct {
	value interface{}
}

func (j *job) Id() string {
	return ""
}

func (j *job) Payload() interface{} {
	return j.value
}

type result struct {
	job    cargo.Job
	result int
}

func (r *result) Id() string {
	return ""
}
func (r *result) Value() interface{} {
	return nil
}
func (r *result) Error() error {
	return nil
}

type simpleHandler struct{}

func (h *simpleHandler) Handle(result cargo.Result) {
	log.Printf("Handle: %+v", result)
}

type worker struct{}

func (w *worker) Process(job cargo.Job) cargo.Result {
	wait := rand.Intn(1000)
	time.Sleep(time.Duration(wait) * time.Millisecond)
	return &result{job, wait}
}

func ExamplePoolHandlerChannel() {
	p, err := cargo.NewPool(100,
		&simpleHandler{},
		func() cargo.Worker {
			return &worker{}
		},
	)
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < 100; i++ {
		p.Execute(&job{})
	}

	p.Close()
}

func TestPoolHandlerChannel(t *testing.T) {
	ExamplePoolHandlerChannel()
}
