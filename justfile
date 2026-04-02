# Build the Go binary in bin
build:
    mkdir -p bin
    go build -o bin/radFS ./cmd/radFS

# Run the filesystem (Usage: just run <folder_name>)
run mnt:
    mkdir -p {{mnt}}
    go run ./cmd/radFS {{mnt}}

# Clean up bin
clean:
    rm -rf bin

# Force unmount if the app crashes (Very helpful for FUSE)
unmount mnt:
    fusermount -u {{mnt}} || umount {{mnt}}

# Format and check Go code logic
check:
    go fmt ./...
    go vet ./...
# Run all tests
test:
    go test -v ./internal/fs/...
