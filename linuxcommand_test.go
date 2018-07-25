package gocommand

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestLinuxExec(t *testing.T) {
	if runtime.GOOS == "linux" {
		var cmd = NewLinuxCommand()
		pid, out, err := cmd.Exec("ls /")

		if err != nil {
			t.Errorf("exec err: %s", err)
		}

		if pid == 0 {
			t.Errorf("exec err: pid is %d", pid)
		}

		if !strings.Contains(fmt.Sprintf("%s", out), "usr") {
			t.Errorf("exec err: [ls /] command not contains usr")
		}
	}
}

func TestLinuxExecAsync(t *testing.T) {
	if runtime.GOOS == "linux" {
		var cmd = NewLinuxCommand()

		rc := make(chan string, 1)
		pid := cmd.ExecAsync(rc, "ls /")

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

		if !strings.Contains(r, "usr") {
			t.Errorf("exec async err: [ls /] command not contains usr")
		}
	}
}

func TestLinuxExecNoWait(t *testing.T) {
	if runtime.GOOS == "linux" {
		var cmd = NewLinuxCommand()
		err := cmd.ExecNoWait("ls /")
		if err != nil {
			t.Errorf("exec nowait err: %s", err)
		}
	}
}