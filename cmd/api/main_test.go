package main

import (
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestDummy")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	err := cmd.Start()
	if err != nil {
		t.Fatalf("Failed to start process: %v", err)
	}

	time.Sleep(2 * time.Second)

	if err := cmd.Process.Signal(syscall.SIGTERM); err != nil {
		t.Fatalf("Failed to send SIGTERM: %v", err)
	}

	state, err := cmd.Process.Wait()
	if err != nil {
		t.Fatalf("Process wait error: %v", err)
	}

	if !state.Success() {
		t.Fatalf("Process exited with non-zero status: %v", state.ExitCode())
	}
}
