//go:build linux

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// `CLONE_NEWUTS": Isolates the UTS (UNIX Timesharing System) namespace,
// which means the new process can have its own hostname.
// `CLONE_NEWPID": Isolates the PID namespace, allowing the child process to
// have a new process ID space, making it seem as if it's the only process
// running inside this namespace.
// `CLONE_NEWNS": Isolates the mount namespace, so that the child process can have
// its own view of the filesystem.
func setSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
}

// Sets the hostname of the isolated environment
func setHostName() error {
	return syscall.Sethostname([]byte("container"))
}

// Isolates the filesystem by using chroot and chdir to change the root to a specific directory
func isolateFilesystem() error {
	// Change root to the specified directory
	if err := syscall.Chroot("/home/your_name/ubuntufs"); err != nil {
		return fmt.Errorf("failed to chroot: %w", err)
	}

	// Change working directory to the new root
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("failed to chdir: %w", err)
	}

	return nil
}

// Mount the proc filesystem inside the container
func mountProc() error {
	if err := syscall.Mount("proc", "proc", "proc", 0, ""); err != nil {
		return fmt.Errorf("failed to mount proc filesystem: %w", err)
	}
	return nil
}

// Unmount the proc filesystem after the process is done
func unmountProc() error {
	if err := syscall.Unmount("proc", 0); err != nil {
		return fmt.Errorf("failed to unmount proc filesystem: %w", err)
	}
	return nil
}

func cg() error {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	cgroupPath := filepath.Join(pids, "your_name")

	// Create the cgroup directory, ignore error if already exists
	if err := os.Mkdir(cgroupPath, 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create cgroup directory: %w", err)
	}

	// Set the maximum number of processes to 20
	if err := ioutil.WriteFile(filepath.Join(cgroupPath, "pids.max"), []byte("20"), 0700); err != nil {
		return fmt.Errorf("failed to set pids.max: %w", err)
	}

	// Set notify_on_release to 1, so the cgroup is removed after use
	if err := ioutil.WriteFile(filepath.Join(cgroupPath, "notify_on_release"), []byte("1"), 0700); err != nil {
		return fmt.Errorf("failed to set notify_on_release: %w", err)
	}

	// Add the current process to the cgroup
	if err := ioutil.WriteFile(filepath.Join(cgroupPath, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700); err != nil {
		return fmt.Errorf("failed to add process to cgroup: %w", err)
	}
	
	return nil
}