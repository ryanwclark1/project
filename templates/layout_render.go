package templates

import (
    "context"
    "io"
)

// RenderLayout renders the Layout with the provided content and title
func RenderLayout(ctx context.Context, w io.Writer, content Templ, title string) {
    tmpl := Layout(content, title)
    tmpl.Render(ctx, w)
}