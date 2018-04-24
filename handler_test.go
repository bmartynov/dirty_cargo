package cargo

import (
	"testing"
)

type result struct {
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

func TestChannelHandler_Handle(t *testing.T) {
	resultReceived := false
	resultChan := make(chan Result)

	go func() {
		for {
			select {
			case <-resultChan:
				resultReceived = true
			}
		}
	}()

	h := newChannelResultHandler(resultChan)
	h.Handle(&result{})

	if !resultReceived {
		t.Fatal("rChan does not received")
	}
}

func TestCallbackHandler_Handle(t *testing.T) {
	resultReceived := false
	var resultHandler = func(_ Result) {
		resultReceived = true
	}

	h := NewCbHandler(resultHandler)
	h.Handle(&result{})

	if !resultReceived {
		t.Fatal("rChan does not received")
	}
}
