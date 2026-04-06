package fs

import (
	"os"
	"sync"
	"sync/atomic"
	"time"

	"bazil.org/fuse/fs"
)

type FS struct {
	Debug bool
}

func New(debug bool) *FS {
	return &FS{Debug: debug}
}

var inodeCounter uint64 = 2

func nextInode() uint64 {
	return atomic.AddUint64(&inodeCounter, 1)
}

func (f *FS) Root() (fs.Node, error) {
	return &Dir{
		inode: 1,
		Nodes: map[string]fs.Node{
			"hello.txt": &File{
				inode: nextInode(),
				data:  []byte("Hello from radFS!\n"),
				mode:  0o666,
				atime: time.Now(),
				mtime: time.Now(),
				ctime: time.Now(),
			},
		},
		fs:    f,
		atime: time.Now(),
		mtime: time.Now(),
		ctime: time.Now(),
	}

	return root, nil
}

type File struct {
	mu    sync.Mutex
	inode uint64
	data  []byte
	mode  uint32
	atime time.Time // read
	mtime time.Time // write | truncate
	ctime time.Time // metadata (setattr)
	uid   uint32
	gid   uint32
}

type Dir struct {
	mu    sync.Mutex
	inode uint64
	Nodes map[string]fs.Node
	fs    *FS
	atime time.Time
	mtime time.Time
	ctime time.Time
}
