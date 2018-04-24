package cargo


type cBHandler struct {
	cb func(Result)
}

func (h *cBHandler) Handle(result Result) {
	h.cb(result)
}

func NewCbHandler(cb func(Result)) ResultHandler {
	return &cBHandler{cb}
}
