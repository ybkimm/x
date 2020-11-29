package templates

import "bytes"

type TestFS struct{}

func (*TestFS) Open(name string) (File, error) {
	if name == "layout.html" {
		return &TestFile{
			Reader: bytes.NewReader(
				[]byte(`
					{{- define "layout" -}}
						{{ template "content" . }}
					{{- end -}}
				`),
			),
		}, nil
	}
	return &TestFile{
		Reader: bytes.NewReader(
			[]byte(`
				{{- define "content" -}}
					{{ .Name }}: hello, world!
				{{- end -}}
				{{- template "layout" . -}}
			`),
		),
	}, nil
}

type TestFile struct {
	*bytes.Reader
}

var _ File = (*TestFile)(nil)

func (*TestFile) Close() error {
	return nil
}
