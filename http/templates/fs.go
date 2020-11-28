package templates

import (
	"io"
	"os"
	"path"
	"path/filepath"
)

type FS interface {
	Open(name string) (File, error)
}

type File interface {
	io.Reader
	io.Closer
}

type OpenError struct {
	error
}

func (e *OpenError) Unwrap() error {
	return e.error
}

type Dir string

func (d Dir) Open(name string) (File, error) {
	dir := string(d)
	if len(dir) == 0 {
		dir = "."
	}
	fullPath := filepath.FromSlash(path.Join(dir, path.Clean("/"+name)))
	f, err := os.Open(fullPath)
	if err != nil {
		return nil, &OpenError{err}
	}
	return f, nil
}
