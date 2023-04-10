package component

import (
	"context"
	"log"
	"sync"

	"github.com/base-org/pessimism/internal/core"
)

// TransformFunc ... Generic transformation function
type TranformFunc func(data core.TransitData) ([]core.TransitData, error)

// Pipe ... Component used to represent any arbitrary computation; pipes must always read from an existing component
// E.G, (ORACLE || CONVEYOR || PIPE) -> PIPE

type Pipe struct {
	ctx    context.Context
	inType core.RegisterType

	tform TranformFunc

	*metaData
}

// NewPipe ... Initializer
func NewPipe(ctx context.Context, tform TranformFunc, inType core.RegisterType, outType core.RegisterType, opts ...Option) (Component, error) {
	// TODO - Validate inTypes size

	router, err := newRouter()
	if err != nil {
		return nil, err
	}

	pipe := &Pipe{
		ctx:    ctx,
		tform:  tform,
		inType: inType,

		metaData: &metaData{
			id:      core.NilCompID(),
			cType:   core.Pipe,
			router:  router,
			ingress: newIngress(),
			state:   Inactive,
			output:  outType,
			RWMutex: &sync.RWMutex{},
		},
	}

	if err := pipe.createEntryPoint(outType); err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(pipe.metaData)
	}

	log.Printf("[%s] Constructed component", pipe.metaData.id.String())
	return pipe, nil
}

// EventLoop ... Driver loop for component that actively subscribes
// to an input channel where transit data is read, transformed, and transitte
// to downstream components
func (p *Pipe) EventLoop() error {
	p.RWMutex.Lock()
	defer p.RWMutex.Unlock()
	log.Printf("[%s][%s] Starting event loop", p.id, p.cType)

	p.metaData.state = Live
	inChan, err := p.GetEntryPoint(p.inType)
	if err != nil {
		return err
	}

	for {

		select {
		case inputData := <-inChan:
			outputData, err := p.tform(inputData)
			if err != nil {
				// TODO - Introduce prometheus call here
				// TODO - Introduce go standard logger (I,E. zap) debug call
				log.Printf("%e", err)
				continue
			}

			if len(outputData) > 0 {
				log.Printf("[%s][%s] Received output data: %s", p.id, p.cType, outputData[0].Type)
			}

			if err := p.router.TransitOutputs(outputData); err != nil {
				log.Printf(transitErr, p.id, p.cType, err.Error())
			}

		// Manager is telling us to shutdown
		case <-p.ctx.Done():
			return nil
		}
	}
}
