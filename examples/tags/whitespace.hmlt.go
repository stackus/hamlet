// Code generated by hamlet - DO NOT EDIT.

package tags

import (
	"bytes"
	"context"
	"github.com/stackus/hamlet"
	"io"
)

// Whitespace will be added between tags and between tags and text
// if that text is on a new line after the tag.
// Whitespace can have subtle effects on the final output of the
// rendered HTML, so it is important to understand how it works.

func Whitespace() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<p>This text has no whitespace between it and the tag.</p>\n<p>\nThis text has whitespace between it and the tag.\n<p>This tag has whitespace between it and the tag above.</p>\n</p>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}

// You can control the whitespace that will be rendered between tags
// in the final output by using the `>` and <` operators.
// The `>` operator will remove all whitespace outside the tag, and
// the `<` operator will remove all whitespace inside the tag.
// These operators must be placed at the end of the tag but before
// either the `!` or `=` operators.
// You can use either or both of these operators on a tag, when using
// both, the order does not matter.

func RemoveWhitespace() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<p>~☢<\nThis text has no whitespace between it and the parent tag.\n>☢~</p>\n<p>\nThere is whitespace between this text and the parent tag.\n>☢~<p>~☢<\nThis text has no whitespace between it and the parent tag.\nThere is also no whitespace between this tag and the sibling text above it.\nFinally, the tag has no whitespace between it and the outer tag.\n>☢~</p>~☢<</p>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}
