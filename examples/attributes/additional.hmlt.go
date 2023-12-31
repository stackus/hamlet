// Code generated by hamlet - DO NOT EDIT.

package main

import (
	"bytes"
	"context"
	"github.com/stackus/hamlet"
	"io"
)

// Additional attributes may be added to an element using the `@attributes`
// command. This command accepts a list of additional attributes in the
// following formats:
// - A map[string]bool - Any attribute with a true value is added as is.
// - A map[string]string - Any attribute with a non-empty string value is
//   added as is.

var boolAttrs = map[string]bool{
	"disabled": true,
	"checked":  false,
}
var strAttrs = map[string]string{
	"type":  "checkbox",
	"value": "foo",
}

func AttributesCmd() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<input"); __err != nil {
			return
		}
		var __var1 string
		__var1, __err = hamlet.BuildAttributeList(boolAttrs, strAttrs)
		if __err != nil {
			return
		}
		if _, __err = __buf.WriteString(__var1); __err != nil {
			return
		}
		if _, __err = __buf.WriteString(">"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}
