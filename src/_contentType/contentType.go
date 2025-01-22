package _contentType

import (
	"path"
	"strings"
)

func Get(filepath string) string {
	switch strings.ToLower(path.Ext(filepath)) {
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	case ".ico":
		return "image/x-icon"
	case ".jpe", ".jpeg":
		return "image/jpeg"
	case ".webp":
		return "image/webp"
	default:
		return "text/plain"
	}
}
