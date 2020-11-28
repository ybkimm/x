package templates

import (
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tpls := New(&TestFS{})
	t.Run("render", func(t *testing.T) {
		var buf strings.Builder
		err := tpls.Render(&buf, "world.html", map[string]string{"Name": "chieri"})
		if err != nil {
			t.Fatalf("render failed: %v", err)
		}
		if buf.String() != "chieri: hello from world.html" {
			t.Fatalf("expected: 'chieri: hello from world.html', got: '%s'", buf.String())
		}
	})
}
