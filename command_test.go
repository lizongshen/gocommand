package gocommand

import (
	"testing"
)

// 测试Command初始化
func TestNewCommand(t *testing.T) {
	var cmd = NewCommand()
	if _, ok := cmd.(Commander); !ok {
		t.Errorf("new Commander err")
	}
}
