package cargo

import "sync"

type batchHandler struct {
	sync.Mutex
	batch []Result
	cb    func([]Result)
}

func (h *batchHandler) reset() {
	h.Lock()
	h.batch = make([]Result, 0, cap(h.batch))
	h.Unlock()
}

func (h *batchHandler) Handle(result Result) {
	h.Lock()
	if len(h.batch) == cap(h.batch) {
		h.cb(h.batch)
		h.reset()
	}
	h.batch = append(h.batch, result)
	h.Unlock()
}

func NewBatchHandler(
	size int,
	handler func([]Result),
) ResultHandler {
	return &batchHandler{
		cb:    handler,
		batch: make([]Result, 0, size),
	}
}
