package cargo

type channelResultHandler struct {
	rChan chan Result
}

func (h *channelResultHandler) Handle(result Result) {
	h.rChan <- result
}

func newChannelResultHandler(rChan chan Result) ResultHandler {
	return &channelResultHandler{rChan}
}

type channelSplittedResultHandler struct {
	rChan    chan Result
	rChanErr chan Result
}

func (h *channelSplittedResultHandler) Handle(result Result) {
	if result.Error() == nil {
		h.rChan <- result
		return
	}
	h.rChanErr <- result
}

func newChannelSplittedResultHandler(rChan, rChanErr chan Result) ResultHandler {
	return &channelSplittedResultHandler{
		rChan,
		rChanErr,
	}
}
