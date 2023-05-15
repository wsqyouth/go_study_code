package loglevel_test

import (
	"go.uber.org/zap"
	"testing"
)

func TestLogLevelFunctions(t *testing.T) {
	testCases := []struct {
		name     string
		function func(string) bool
		input    string
		expected bool
	}{
		{"Debug", loglevel.IsDebug, "debug", true},
		{"Debug", loglevel.IsDebug, "info", false},
		{"Info", loglevel.IsInfo, "info", true},
		{"Info", loglevel.IsInfo, "error", false},
		{"Warn", loglevel.IsWarn, "warn", true},
		{"Warn", loglevel.IsWarn, "debug", false},
		{"Error", loglevel.IsError, "error", true},
		{"Error", loglevel.IsError, "info", false},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.function(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
