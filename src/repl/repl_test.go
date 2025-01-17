package repl_test

import (
	"interpreter/src/repl"
	"io"
	"testing"
)

func TestStart(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		in  io.Reader
		out io.Writer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repl.Start(tt.in, tt.out)
		})
	}
}
