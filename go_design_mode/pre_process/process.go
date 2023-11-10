package main

import "context"

// Checker
type Checker interface {
	Check(ctx context.Context) (err error)
}

type CheckerFunc func(ctx context.Context) (err error)

func (f CheckerFunc) Check(ctx context.Context) (err error) {
	return f(ctx)
}

type Processor interface {
	Process(ctx context.Context) (err error)
}

// ProcessorFunc
type ProcessorFunc func(ctx context.Context) (err error)

// Process
func (f ProcessorFunc) Process(ctx context.Context) (err error) {
	return f(ctx)
}

type PreProcessor struct {
	checkers   []Checker
	processors []Processor
}

func NewPreProcessor(checkers []Checker, processors []Processor) *PreProcessor {
	return &PreProcessor{
		checkers:   checkers,
		processors: processors,
	}
}

func (p *PreProcessor) Do(ctx context.Context) (err error) {
	for _, v := range p.checkers {
		if err := v.Check(ctx); err != nil {
			return err
		}
	}

	for _, t := range p.processors {
		if err := t.Process(ctx); err != nil {
			return err
		}
	}

	return nil
}
