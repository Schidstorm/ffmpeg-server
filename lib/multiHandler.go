package lib

type MultiHandler struct {
	handlers            []TaskHandler
	currentHandlerIndex int
}

func (h *MultiHandler) Progress() float32 {
	divider := float32(1) / float32(len(h.handlers))
	return float32(h.currentHandlerIndex)*divider + h.handlers[h.currentHandlerIndex].Progress()/divider
}

func (h *MultiHandler) Run() error {
	for index, handler := range h.handlers {
		h.currentHandlerIndex = index
		err := handler.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMultiHandler(handlers ...TaskHandler) *MultiHandler {
	return &MultiHandler{
		handlers:            handlers,
		currentHandlerIndex: 0,
	}
}
