package web

import "embed"

//go:embed "assets"
var staticFiles embed.FS
