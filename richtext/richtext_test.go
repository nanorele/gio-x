package richtext

import (
	"image"
	"testing"
	"time"

	"github.com/uorg-saver/gio/font/gofont"
	"github.com/uorg-saver/gio/io/input"
	"github.com/uorg-saver/gio/layout"
	"github.com/uorg-saver/gio/op"
	"github.com/uorg-saver/gio/text"
	"github.com/uorg-saver/gio/unit"
	"github.com/uorg-saver/gio/widget/material"
)

// TestNilInteractiveText ensures that it is safe to lay out
// richtext with a nil state when none of the spans are
// interactive.
func TestNilInteractiveText(t *testing.T) {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	spans := []SpanStyle{
		{
			Size:    12,
			Content: "Hello",
		},
		{
			Size:    12,
			Content: "world",
		},
	}
	var ops op.Ops
	gtx := layout.Context{
		Constraints: layout.Exact(image.Pt(100, 100)),
		Metric: unit.Metric{
			PxPerDp: 1,
			PxPerSp: 1,
		},
		Source: input.Source{},
		Now:    time.Now(),
		Ops:    &ops,
	}

	Text(nil, th.Shaper, spans...).Layout(gtx)
}
