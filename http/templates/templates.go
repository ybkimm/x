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

func (t *Templates) UseFuncs(funcs template.FuncMap) {
	t.tpl.Funcs(funcs)
}

func (t *Templates) parse(tpl *template.Template, name, src string) (*template.Template, error) {
	if tpl == nil {
		clone, err := t.tpl.Clone()
		if err != nil {
			return nil, err
		}
		tpl = clone
	}
	return tpl.New(name).Parse(src)
}

// UseTemplate loads a template file to base.
func (t *Templates) UseTemplate(name string) error {
	buf, err := t.open(name)
	if err != nil {
		return err
	}

	tpl, err := t.parse(t.tpl, "name", string(buf))
	if err != nil {
		return err
	}

	t.tpl = tpl
	return nil
}

func (t *Templates) render(w io.Writer, name string, data interface{}) error {
	buf, err := t.open(name)
	if err != nil {
		return err
	}

	tpl, err := t.parse(nil, name, string(buf))
	if err != nil {
		return err
	}

	return tpl.ExecuteTemplate(w, name, data)
}

// Render renders the template and write it's result to w.
func (t *Templates) Render(w io.Writer, name string, data interface{}) error {
	return t.render(w, name, data)
}
