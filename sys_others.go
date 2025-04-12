//go:build !linux

package main

import (
	"errors"
	"os/exec"
)

func setSysProcAttr(cmd *exec.Cmd) {
	// No-op on non-Linux systems
}

func setHostName() error {
	return errors.New("setting hostname is not supported on this platform")
}

// Filesystem isolation is not supported on non-Linux systems
func isolateFilesystem() error {
	return errors.New("filesystem isolation (chroot) is not supported on this platform")
}

// Proc filesystem mounting is not supported on non-Linux systems
func mountProc() error {
	return errors.New("mounting proc filesystem is not supported on this platform")
}

// Proc filesystem unmounting is not supported on non-Linux systems
func unmountProc() error {
	return errors.New("unmounting proc filesystem is not supported on this platform")
}

func cg() error {
	return errors.New("cgroups are not supported on this platform")
}