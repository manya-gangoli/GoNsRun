# 📦 Container From Scratch

A low-level container runtime built from first principles in Go. This project peels back the layers of abstraction to reveal how container technology actually works.

## What & Why

I built this project to deeply understand the core technologies behind containers. By implementing containerization primitives directly with Linux kernel features, I've learned that containers aren't magic - they're just clever applications of:

- Process namespaces for isolation
- Control groups for resource limits
- Filesystem manipulation for environment consistency

## Implemented Features

- [x] **Process Isolation**: Using Linux namespaces (`CLONE_NEWUTS`, `CLONE_NEWPID`, `CLONE_NEWNS`)
- [x] **Custom Environment**: Isolated hostname and process tree
- [x] **Filesystem Isolation**: Restricted view using `chroot`
- [x] **Process Management**: `/proc` filesystem mounting for process visibility
- [x] **Resource Controls**: Basic limits using cgroups
- [x] **Cross-Platform**: Graceful fallbacks for non-Linux systems

## Future Enhancements

- [ ] Network namespace isolation for container networking
- [ ] User namespace implementation for improved security
- [ ] Volume mounting for data persistence
- [ ] Performance optimizations

## 🛠️ How It Works

This project demonstrates containerization from scratch using these key mechanisms:

### 1. Process Isolation

```go
// Create isolated process environment using namespaces
cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
}
```

### 2. Filesystem Boundaries

```go
// Change root filesystem to create isolation
syscall.Chroot("/path/to/container/root")
syscall.Chdir("/")
```

### 3. Resource Management

```go
// Limit container to 20 processes max
ioutil.WriteFile(filepath.Join(pids, "container/pids.max"), []byte("20"), 0700)
```

## Usage

### For Linux Users:

```bash
# Run a bash shell inside a container
go run main.go run /bin/bash
```

### For Everyone Else (Docker):

```bash
# Build and run using Docker
docker build -t container-from-scratch .
docker run --rm -it --privileged container-from-scratch ./main run /bin/bash
```

> **Note:** The `--privileged` flag is needed because we're running container tech inside a container.

## Learn More

Check the `docs` folder for detailed explanations of the concepts behind this implementation. This project is for educational purposes - not production use!

## Technical Details

This implementation uses:

- Go's syscall package for direct kernel interaction
- Linux namespaces for process isolation
- cgroups for resource management
- chroot for filesystem isolation
- proc filesystem for process visibility
