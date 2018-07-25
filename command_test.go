package gocommand

import (
	"testing"
)

func TestNewCommand(t *testing.T) {
	var cmd = NewCommand()
	if _, ok := cmd.(Commander); !ok {
		t.Errorf("new Commander err")
	}
}
