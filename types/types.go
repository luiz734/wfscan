package types

import (
	"fmt"
	"path/filepath"

	"github.com/dustin/go-humanize"
)

type FileInfo struct {
	Dir  string
	Name string
	Size uint64
}

func NewFileInfo(dir, name string, size int64) *FileInfo {
	return &FileInfo{
		Dir:  dir,
		Name: name,
		Size: uint64(size),
	}
}

func (f FileInfo) String() string {
	readableSize := humanize.Bytes(f.Size)
	fullPath := filepath.Join(f.Dir, f.Name)
	return fmt.Sprintf("%s:\t%s", readableSize, fullPath)
}
