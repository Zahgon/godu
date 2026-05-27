package files

import (
	"os"
	"sync"
)

// File structure representing files and folders with their accumulated sizes
type File struct {
	Name   string
	Parent *File
	Size   int64
	IsDir  bool
	Files  []*File
}

// Path builds a file system location for given file
func (f *File) Path() string { _ = "STUB: not implemented"; return "" }

// UpdateSize goes through subfiles and subfolders and accumulates their size
func (f *File) UpdateSize() { _ = "STUB: not implemented"; return }

// ReadDir function can return list of files for given folder path
type ReadDir func(dirname string) ([]os.FileInfo, error)

// ShouldIgnoreFolder function decides whether a folder should be ignored
type ShouldIgnoreFolder func(absolutePath string) bool

func ignoringReadDir(shouldIgnore ShouldIgnoreFolder, originalReadDir ReadDir) ReadDir {
	_ = "STUB: not implemented"
	return *new(ReadDir)
}

// WalkFolder will go through a given folder and subfolders and produces file structure
// with aggregated file sizes
func WalkFolder(
	path string,
	readDir ReadDir,
	ignoreFunction ShouldIgnoreFolder,
	progress chan<- int,
) *File {
	_ = "STUB: not implemented"
	return nil
}

func walkSubFolderConcurrently(
	path string,
	parent *File,
	readDir ReadDir,
	c chan bool,
	wg *sync.WaitGroup,
	progress chan<- int,
) *File {
	_ = "STUB: not implemented"
	return nil
}

// Root dir
// TODO unit test this Join

func updateProgress(progress chan<- int, count *int) { _ = "STUB: not implemented"; return }
