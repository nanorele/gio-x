package richtext

import (
	"image"
	"testing"
	"time"

	"github.com/nanorele/gio/font/gofont"
	"github.com/nanorele/gio/io/input"
	"github.com/nanorele/gio/layout"
	"github.com/nanorele/gio/op"
	"github.com/nanorele/gio/text"
	"github.com/nanorele/gio/unit"
	"github.com/nanorele/gio/widget/material"
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
