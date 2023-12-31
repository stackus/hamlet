// Code generated by hamlet - DO NOT EDIT.

package tags

import (
	"bytes"
	"context"
	"github.com/stackus/hamlet"
	"io"
)

// You may specify the type of tag that you want created with `%`.

func SpecifyTag() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<p>This is a paragraph tag.</p>\n<main>This is a main tag.</main>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}

// You may also let the tag default to a `div` when using the id or
// class syntax's, `#` and `.` respectively.

func DefaultToDivs() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<div id=\"main\">This is a div tag with an id of `main`.</div>\n<div class=\"main\">This is a div tag with a class of `main`.</div>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}

// The three may also be combined. The `%` must come first, followed
// by either the `#` or `.`. The `#` and `.` may be in any order.

func Combined() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<p id=\"main\">This is a paragraph tag with an id of `main`.</p>\n<main class=\"main\">This is a main tag with a class of `main`.</main>\n<div id=\"main\" class=\"main\">This is a div tag with an id and class of `main`.</div>\n<p id=\"main\" class=\"main\">This is a paragraph tag with an id and class of `main`.</p>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}

// The class operator may be repeated to add multiple classes.
// Repeating the id operator will result in the id being overwritten
// but will not throw an error.

func MultipleClasses() hamlet.Template {
	return hamlet.TemplateFunc(func(ctx context.Context, __w io.Writer) (__err error) {
		__buf, __isBuf := __w.(*bytes.Buffer)
		if !__isBuf {
			__buf = hamlet.GetBuffer()
			defer hamlet.ReleaseBuffer(__buf)
		}
		var __children hamlet.Template
		ctx, __children = hamlet.PopChildren(ctx)
		_ = __children
		if _, __err = __buf.WriteString("<div class=\"main main2\">This is a div tag with two classes, `main` and `main2`.</div>\n<div id=\"main2\">This is a div tag with an id of `main2`.</div>\n"); __err != nil {
			return
		}
		if !__isBuf {
			_, __err = __w.Write(hamlet.NukeWhitespace(__buf.Bytes()))
		}
		return
	})
}
