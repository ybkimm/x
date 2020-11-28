package templates

import (
	"html/template"
	"io"
	"io/ioutil"
)

// Templates hold few instances for render templates.
type Templates struct {
	fs  FS
	tpl *template.Template
}

// New returns new Templates instance.
func New(fs FS) *Templates {
	return &Templates{
		fs:  fs,
		tpl: template.New("root"),
	}
}

func (t *Templates) open(name string) ([]byte, error) {
	f, err := t.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func (t *Templates) parse(name, src string) (*template.Template, error) {
	clone, err := t.tpl.Clone()
	if err != nil {
		return nil, err
	}
	return clone.New(name).Parse(src)
}

// Render renders the template and write it's result to w.
func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	buf, err := t.open(name)
	if err != nil {
		return err
	}

	tpl, err := t.parse(name, string(buf))
	if err != nil {
		return err
	}

	return tpl.ExecuteTemplate(w, name, data)
}
