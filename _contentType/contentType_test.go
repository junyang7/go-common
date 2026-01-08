package _contentType

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestGet_Text(t *testing.T) {
	// HTML
	{
		_assert.Equal(t, "text/html", Get("index.html"))
		_assert.Equal(t, "text/html", Get("page.htm"))
		_assert.Equal(t, "text/html", Get("INDEX.HTML")) // 大小写
	}

	// CSS
	{
		_assert.Equal(t, "text/css", Get("style.css"))
	}

	// JavaScript
	{
		_assert.Equal(t, "text/javascript", Get("app.js"))
	}

	// XML
	{
		_assert.Equal(t, "text/xml", Get("data.xml"))
	}

	// CSV
	{
		_assert.Equal(t, "text/csv", Get("data.csv"))
	}

	// TXT
	{
		_assert.Equal(t, "text/plain", Get("readme.txt"))
	}

	// Markdown
	{
		_assert.Equal(t, "text/markdown", Get("README.md"))
	}
}

func TestGet_Image(t *testing.T) {
	// JPEG
	{
		_assert.Equal(t, "image/jpeg", Get("photo.jpg"))
		_assert.Equal(t, "image/jpeg", Get("photo.jpeg"))
		_assert.Equal(t, "image/jpeg", Get("photo.jpe"))
	}

	// PNG
	{
		_assert.Equal(t, "image/png", Get("icon.png"))
	}

	// GIF
	{
		_assert.Equal(t, "image/gif", Get("animation.gif"))
	}

	// WebP
	{
		_assert.Equal(t, "image/webp", Get("image.webp"))
	}

	// SVG
	{
		_assert.Equal(t, "image/svg+xml", Get("logo.svg"))
	}

	// BMP
	{
		_assert.Equal(t, "image/bmp", Get("bitmap.bmp"))
	}

	// ICO
	{
		_assert.Equal(t, "image/x-icon", Get("favicon.ico"))
	}
}

func TestGet_Audio(t *testing.T) {
	// MP3
	{
		_assert.Equal(t, "audio/mpeg", Get("song.mp3"))
	}

	// WAV
	{
		_assert.Equal(t, "audio/wav", Get("sound.wav"))
	}

	// OGG
	{
		_assert.Equal(t, "audio/ogg", Get("audio.ogg"))
		_assert.Equal(t, "audio/ogg", Get("audio.opus"))
	}

	// FLAC
	{
		_assert.Equal(t, "audio/flac", Get("music.flac"))
	}

	// AAC
	{
		_assert.Equal(t, "audio/aac", Get("track.aac"))
	}
}

func TestGet_Video(t *testing.T) {
	// MP4
	{
		_assert.Equal(t, "video/mp4", Get("video.mp4"))
	}

	// WebM
	{
		_assert.Equal(t, "video/webm", Get("clip.webm"))
	}

	// OGV
	{
		_assert.Equal(t, "video/ogg", Get("movie.ogv"))
	}

	// AVI
	{
		_assert.Equal(t, "video/avi", Get("film.avi"))
	}

	// MPEG
	{
		_assert.Equal(t, "video/mpeg", Get("video.mpeg"))
		_assert.Equal(t, "video/mpeg", Get("video.mpg"))
	}

	// MOV
	{
		_assert.Equal(t, "video/quicktime", Get("movie.mov"))
	}
}

func TestGet_Application(t *testing.T) {
	// JSON
	{
		_assert.Equal(t, "application/json", Get("data.json"))
	}

	// PDF
	{
		_assert.Equal(t, "application/pdf", Get("document.pdf"))
	}

	// ZIP
	{
		_assert.Equal(t, "application/zip", Get("archive.zip"))
	}

	// TAR
	{
		_assert.Equal(t, "application/x-tar", Get("backup.tar"))
	}

	// GZIP
	{
		_assert.Equal(t, "application/gzip", Get("compressed.gz"))
	}

	// 7Z
	{
		_assert.Equal(t, "application/x-7z-compressed", Get("archive.7z"))
	}

	// RAR
	{
		_assert.Equal(t, "application/vnd.rar", Get("archive.rar"))
	}

	// YAML
	{
		_assert.Equal(t, "application/x-yaml", Get("config.yaml"))
		_assert.Equal(t, "application/x-yaml", Get("config.yml"))
	}

	// TOML
	{
		_assert.Equal(t, "application/toml", Get("config.toml"))
	}
}

func TestGet_Font(t *testing.T) {
	// WOFF
	{
		_assert.Equal(t, "font/woff", Get("font.woff"))
	}

	// WOFF2
	{
		_assert.Equal(t, "font/woff2", Get("font.woff2"))
	}

	// TTF
	{
		_assert.Equal(t, "font/ttf", Get("font.ttf"))
	}

	// OTF
	{
		_assert.Equal(t, "font/otf", Get("font.otf"))
	}
}

func TestGet_Script(t *testing.T) {
	// TypeScript
	{
		_assert.Equal(t, "text/typescript", Get("app.ts"))
	}

	// JSX
	{
		_assert.Equal(t, "text/jsx", Get("component.jsx"))
	}

	// TSX
	{
		_assert.Equal(t, "text/tsx", Get("component.tsx"))
	}
}

func TestGet_EdgeCases(t *testing.T) {
	// 无扩展名
	{
		_assert.Equal(t, "application/octet-stream", Get("file"))
	}

	// 空字符串
	{
		_assert.Equal(t, "application/octet-stream", Get(""))
	}

	// 未知扩展名
	{
		_assert.Equal(t, "application/octet-stream", Get("file.unknown"))
	}

	// 大小写混合
	{
		_assert.Equal(t, "image/jpeg", Get("photo.JPG"))
		_assert.Equal(t, "image/png", Get("image.PNG"))
		_assert.Equal(t, "text/html", Get("index.HTML"))
	}

	// 路径格式
	{
		_assert.Equal(t, "image/jpeg", Get("/path/to/image.jpg"))
		_assert.Equal(t, "text/css", Get("../styles/main.css"))
		_assert.Equal(t, "text/javascript", Get("./js/app.js"))
	}

	// 多个点号
	{
		_assert.Equal(t, "image/jpeg", Get("file.backup.jpg"))
		_assert.Equal(t, "application/gzip", Get("archive.tar.gz"))
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("image.jpg")
	}
}

func BenchmarkGet_Unknown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("file.unknown")
	}
}

