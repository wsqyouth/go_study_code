package main

import (
	"context"
	"errors"
	"testing"
)

type MockChecker struct {
	Valid bool
}

func (v *MockChecker) Check(ctx context.Context) error {
	if v.Valid {
		return nil
	}
	return errors.New("check error")
}

type MockProcess struct{}

func (t *MockProcess) Process(ctx context.Context) error {
	return nil
}

func TestPreProcessorALLCondition(t *testing.T) {
	tests := []struct {
		name       string
		isCheckErr bool
		wantErr    bool
	}{
		{
			name:    "normal test",
			wantErr: false,
		},
		{
			name:       "err check test",
			isCheckErr: true,
			wantErr:    true,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// Create a PreProcessor instance with mock Checker and Processs
			p := &PreProcessor{
				checkers:   []Checker{&MockChecker{Valid: true}},
				processors: []Processor{&MockProcess{}},
			}
			if tt.isCheckErr {
				p = &PreProcessor{
					checkers:   []Checker{&MockChecker{Valid: false}},
					processors: []Processor{&MockProcess{}},
				}
			}
			if err := p.Do(ctx); (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
