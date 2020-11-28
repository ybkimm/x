package templates

import "bytes"

type TestFS struct{}

func (*TestFS) Open(name string) (File, error) {
	return &TestFile{
		Reader: bytes.NewReader(append([]byte("{{ .Name }}: hello from "), name...)),
	}, nil
}

type TestFile struct {
	*bytes.Reader
}

var _ File = (*TestFile)(nil)

func (*TestFile) Close() error {
	return nil
}
