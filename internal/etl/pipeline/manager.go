package pipeline

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/base-org/pessimism/internal/config"
	"github.com/base-org/pessimism/internal/core"
	"github.com/base-org/pessimism/internal/etl/component"
	"github.com/base-org/pessimism/internal/etl/registry"
)

type Manager struct {
	ctx context.Context

	dag       *cGraph
	pipeLines map[core.PipelineID]PipeLine
	wg        *sync.WaitGroup
}

func NewManager(ctx context.Context) *Manager {
	return &Manager{
		ctx:       ctx,
		dag:       newGraph(),
		pipeLines: make(map[core.PipelineID]PipeLine, 0),
		wg:        &sync.WaitGroup{},
	}
}

func (manager *Manager) CreateRegisterPipeline(ctx context.Context, cfg *config.PipelineConfig) (core.PipelineID, error) {
	log.Printf("Constructing register pipeline for %s", cfg.DataType)

	register, err := registry.GetRegister(cfg.DataType)
	if err != nil {
		return core.NilPipelineID(), err
	}

	components := make([]component.Component, 0)
	registers := append([]*core.DataRegister{register}, register.Dependencies...)
	log.Printf("%+v", registers)

	var prevID core.ComponentID = core.NilCompID()
	lastReg := registers[len(registers)-1]

	pID := core.MakePipelineID(
		cfg.PipelineType,
		core.MakeComponentID(cfg.PipelineType, registers[0].ComponentType, registers[0].DataType, cfg.Network),
		core.MakeComponentID(cfg.PipelineType, lastReg.ComponentType, lastReg.DataType, cfg.Network),
	)

	for i, register := range registers {
		// NOTE - This doesn't consider the circumstance where a requested pipeline already exists but requires some backfill to run
		cID := core.MakeComponentID(cfg.PipelineType, register.ComponentType, register.DataType, cfg.Network)
		if err != nil {
			return core.NilPipelineID(), err
		}

		if !manager.dag.componentExists(cID) {

			comp, err := inferComponent(ctx, cfg, cID, register)
			if err != nil {
				return core.NilPipelineID(), err
			}

			manager.dag.addComponent(cID, comp)
		}

		component, err := manager.dag.getComponent(cID)
		if err != nil {
			return core.NilPipelineID(), err
		}

		if i != 0 { // IE we've passed the pipeline's origin node
			if err := manager.dag.addEdge(cID, prevID); err != nil {
				return core.NilPipelineID(), err
			}
		}

		prevID = component.ID()
		components = append(components, component)
	}

	pipeLine, err := NewPipeLine(pID, components)
	if err != nil {
		return core.NilPipelineID(), err
	}

	manager.pipeLines[pID] = pipeLine
	// TODO - Update pipeline entries with component entry struct within componentMap

	return pID, nil
}

func (manager *Manager) RunPipeline(id core.PipelineID) error {

	pipeLine, found := manager.pipeLines[id]
	if !found {
		return fmt.Errorf("Could not find pipeline for id: %s", id)
	}

	log.Printf("[%s] Running pipeline: %s ", pipeLine.ID().String(), pipeLine.String())
	return pipeLine.RunPipeline(manager.wg)
}

func (manager *Manager) AddPipelineDirective(pID core.PipelineID, cID core.ComponentID, outChan chan core.TransitData) error {
	pipeLine, found := manager.pipeLines[pID]
	if !found {
		return fmt.Errorf("Could not find pipeline for id: %s", pID)
	}

	return pipeLine.AddDirective(cID, outChan)
}

func inferComponent(ctx context.Context, cfg *config.PipelineConfig, id core.ComponentID,
	register *core.DataRegister) (component.Component, error) {
	log.Printf("Constructing %s component for register %s", register.ComponentType, register.DataType)

	switch register.ComponentType {
	case core.Oracle:
		init, success := register.ComponentConstructor.(component.OracleConstructorFunc)
		if !success {
			return nil, fmt.Errorf("could not cast constructor to oracle constructor type")
		}

		// NOTE ... We assume at most 1 oracle per register pipeline
		return init(ctx, cfg.PipelineType, cfg.OracleCfg, component.WithID(id))

	case core.Pipe:
		init, success := register.ComponentConstructor.(component.PipeConstructorFunc)
		if !success {
			return nil, fmt.Errorf("could not cast constructor to pipe constructor type")
		}

		return init(ctx, component.WithID(id))

	case core.Aggregator:
		return nil, fmt.Errorf("aggregator component has yet to be implemented")

	default:
		return nil, fmt.Errorf("unknown component type provided")
	}
}
