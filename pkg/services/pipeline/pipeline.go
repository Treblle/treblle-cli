package pipeline

// Handler interface to be implemented on each Pipeline Task
type Handler interface {
	Process(input interface{}) (output interface{}, err error)
}

// Pipeline holds all handlers
type Pipeline struct {
	handlers []Handler
	inChan   chan interface{}
	outChan  chan interface{}
	errChan  chan error
}

// NewPipeline will create a new workflow pipeline
func NewPipeline(handlers ...Handler) *Pipeline {
	return &Pipeline{
		handlers: handlers,
		inChan:   make(chan interface{}),
		outChan:  make(chan interface{}),
		errChan:  make(chan error),
	}
}

// Pipe a new handler into the pipeline.
func (p *Pipeline) Pipe(h Handler) *Pipeline {
	p.handlers = append(p.handlers, h)
	return p
}

// Start running the pipeline
func (p *Pipeline) Start() {
	go func() {
		for input := range p.inChan {
			current := input
			var err error
			for _, handler := range p.handlers {
				current, err = handler.Process(current)
				if err != nil {
					p.errChan <- err
					break
				}
			}
			if err == nil {
				p.outChan <- current
			}
		}
		close(p.outChan)
		close(p.errChan)
	}()
}

func (p *Pipeline) Input() chan<- interface{} {
	return p.inChan
}

func (p *Pipeline) Output() <-chan interface{} {
	return p.outChan
}

func (p *Pipeline) Errors() <-chan error {
	return p.errChan
}
