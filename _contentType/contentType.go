package _contentType

import (
	"path"
	"strings"
)

func Get(filepath string) string {
	switch strings.ToLower(path.Ext(filepath)) {
	// 文本类型
	case ".html", ".htm":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "text/javascript"
	case ".xml":
		return "text/xml"
	case ".csv":
		return "text/csv"
	case ".txt":
		return "text/plain"
	case ".md":
		return "text/markdown"
	// 图片类型
	case ".ico":
		return "image/x-icon"
	case ".jpg", ".jpe", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg":
		return "image/svg+xml"
	case ".bmp":
		return "image/bmp"
	// 音频类型
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".ogg", ".opus":
		return "audio/ogg"
	case ".flac":
		return "audio/flac"
	case ".aac":
		return "audio/aac"
	// 视频类型
	case ".mp4":
		return "video/mp4"
	case ".webm":
		return "video/webm"
	case ".ogv":
		return "video/ogg"
	case ".avi":
		return "video/avi"
	case ".mpeg", ".mpg":
		return "video/mpeg"
	case ".mov":
		return "video/quicktime"
	// 应用类型
	case ".json":
		return "application/json"
	case ".pdf":
		return "application/pdf"
	case ".zip":
		return "application/zip"
	case ".tar":
		return "application/x-tar"
	case ".gz":
		return "application/gzip"
	case ".7z":
		return "application/x-7z-compressed"
	case ".rar":
		return "application/vnd.rar"
	case ".yaml", ".yml":
		return "application/x-yaml"
	case ".toml":
		return "application/toml"
	// 字体类型
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".ttf":
		return "font/ttf"
	case ".otf":
		return "font/otf"
	// 脚本/代码类型
	case ".ts":
		return "text/typescript"
	case ".jsx":
		return "text/jsx"
	case ".tsx":
		return "text/tsx"
	// 默认类型
	default:
		return "application/octet-stream"
	}
}
