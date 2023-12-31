// Code generated by hamlet - DO NOT EDIT.

package main

import (
	"bytes"
	"context"
	"github.com/stackus/hamlet"
	"io"
)

// Attributes may conditionally appear in the output. For example, you may
// want to add the `disabled` attribute to a button if a variable is true.
// When the same variable is false, you want the attribute to be omitted.
// A special operator `?` is used to conditionally add an attribute. It is
// used in place of the colon.
// A dynamic attribute value that evaluates to a boolean must be used with
// the conditional operator. This means that boolean values, not strings,
// are expected. Statements such as `#{true}` or `#{foo == "bar"}` are
// valid too.

var disabled = true

var foo = "bar"

func ConditionalAttrs() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<button"); __err != nil {
			return
		}
		if disabled {
			if _, __err = __buf.WriteString(" disabled"); __err != nil {
				return
			}
		}
		if _, __err = __buf.WriteString(">Click me!</button>\n<button"); __err != nil {
			return
		}
		if foo == "bar" {
			if _, __err = __buf.WriteString(" disabled"); __err != nil {
				return
			}
		}
		if _, __err = __buf.WriteString(">Click me!</button>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}
