package templates

import (
	"strconv"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	tpls := New(&TestFS{})
	tpls.UseTemplate("layout.html")

	t.Run("render", func(t *testing.T) {
		var expected = "chieri: hello, world!"

		var buf strings.Builder
		err := tpls.Render(&buf, "world.html", map[string]string{"Name": "chieri"})
		if err != nil {
			t.Fatalf("render failed: %v", err)
		}
		if buf.String() != expected {
			t.Fatalf("expected: %s, got: %s", strconv.Quote(expected), strconv.Quote(buf.String()))
		}
	})
}
