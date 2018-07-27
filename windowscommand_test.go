package gocommand

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// 测试Windows Exec
func TestWindowsExec(t *testing.T) {
	if runtime.GOOS == "windows" {
		var cmd = NewWindowsCommand()
		pid, out, err := cmd.Exec("dir c:/")

		if err != nil {
			t.Errorf("exec err: %s", err)
		}

		if pid == 0 {
			t.Errorf("exec err: pid is %d", pid)
		}

		if !strings.Contains(fmt.Sprintf("%s", out), "Windows") {
			t.Errorf("exec err: [dir c:/] command not contains Windows")
		}
	}
}

// 测试Windows异步Exec
func TestWindowsExecAsync(t *testing.T) {
	if runtime.GOOS == "windows" {
		var cmd = NewWindowsCommand()

		rc := make(chan string, 1)
		pid := cmd.ExecAsync(rc, "dir c:/")

		r, ok := <-rc
		if !ok {
			t.Errorf("exec async read chan err!")
		}

		if r == "" {
			t.Errorf("exec async err!")
		}

		if pid == 0 {
			t.Errorf("exec async err: pid is %d", pid)
		}

		if !strings.Contains(r, "Windows") {
			t.Errorf("exec async err: [dir c:/ command not contains Windows")
		}
	}
}

// 测试Windows下的Exec(无等待)
func TestExecIgnoreResult(t *testing.T) {
	if runtime.GOOS == "windows" {
		var cmd = NewWindowsCommand()
		err := cmd.ExecIgnoreResult("dir c:/")
		if err != nil {
			t.Errorf("exec nowait err: %s", err)
		}
	}
}
