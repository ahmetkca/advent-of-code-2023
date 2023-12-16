package main

import "context"

// PipelineStage is an interface for pipeline stages
type PipelineStage interface {
	Process(ctx context.Context, in <-chan interface{}) <-chan interface{}
}

// PipelineBuilder is used to build and link pipeline stages
type PipelineBuilder struct {
	stages []PipelineStage
	ctx    context.Context
}

// NewPipelineBuilder creates a new PipelineBuilder with the provided context
func NewPipelineBuilder(ctx context.Context) *PipelineBuilder {
	return &PipelineBuilder{
		ctx: ctx,
	}
}

// AddStage adds a new stage to the pipeline
func (pb *PipelineBuilder) AddStage(stage PipelineStage) *PipelineBuilder {
	pb.stages = append(pb.stages, stage)
	return pb
}

// Build constructs the pipeline and returns the final output channel
func (pb *PipelineBuilder) Build(input <-chan interface{}) <-chan interface{} {
	current := input
	for _, stage := range pb.stages {
		current = stage.Process(pb.ctx, current)
	}
	return current
}
