# _contentType - 内容类型识别

根据文件扩展名获取 MIME 类型（Content-Type）。

---

## 🎯 作用

用于 HTTP 响应、文件上传等场景，根据文件名自动识别正确的 Content-Type。

---

## 💡 使用

```go
import "_contentType"

// 获取 MIME 类型
contentType := _contentType.Get("image.jpg")
// 返回: "image/jpeg"
```

---

## 📖 API

### Get

```go
func Get(filepath string) string
```

根据文件路径（或文件名）获取 Content-Type。

**参数：**
- `filepath`: 文件路径或文件名（如 "image.jpg" 或 "/path/to/file.png"）

**返回：**
- `string`: MIME 类型（如 "image/jpeg"）

**特性：**
- ✅ 自动转换为小写（大小写不敏感）
- ✅ 支持完整路径（只取扩展名）
- ✅ 未知类型返回 "application/octet-stream"

---

## 📋 支持的文件类型

### 文本类型

| 扩展名 | Content-Type |
|--------|--------------|
| .html, .htm | text/html |
| .css | text/css |
| .js | text/javascript |
| .xml | text/xml |
| .csv | text/csv |
| .txt | text/plain |
| .md | text/markdown |

### 图片类型

| 扩展名 | Content-Type |
|--------|--------------|
| .jpg, .jpeg, .jpe | image/jpeg |
| .png | image/png |
| .gif | image/gif |
| .webp | image/webp |
| .svg | image/svg+xml |
| .bmp | image/bmp |
| .ico | image/x-icon |

### 音频类型

| 扩展名 | Content-Type |
|--------|--------------|
| .mp3 | audio/mpeg |
| .wav | audio/wav |
| .ogg, .opus | audio/ogg |
| .flac | audio/flac |
| .aac | audio/aac |

### 视频类型

| 扩展名 | Content-Type |
|--------|--------------|
| .mp4 | video/mp4 |
| .webm | video/webm |
| .ogv | video/ogg |
| .avi | video/avi |
| .mpeg, .mpg | video/mpeg |
| .mov | video/quicktime |

### 应用类型

| 扩展名 | Content-Type |
|--------|--------------|
| .json | application/json |
| .pdf | application/pdf |
| .zip | application/zip |
| .tar | application/x-tar |
| .gz | application/gzip |
| .7z | application/x-7z-compressed |
| .rar | application/vnd.rar |
| .yaml, .yml | application/x-yaml |
| .toml | application/toml |

### 字体类型

| 扩展名 | Content-Type |
|--------|--------------|
| .woff | font/woff |
| .woff2 | font/woff2 |
| .ttf | font/ttf |
| .otf | font/otf |

### 脚本/代码类型

| 扩展名 | Content-Type |
|--------|--------------|
| .ts | text/typescript |
| .jsx | text/jsx |
| .tsx | text/tsx |

### 默认类型

| 场景 | Content-Type |
|------|--------------|
| 无扩展名 | application/octet-stream |
| 未知扩展名 | application/octet-stream |

---

## 🚀 使用示例

### HTTP 文件服务

```go
import (
    "_contentType"
    "net/http"
)

func ServeFile(w http.ResponseWriter, r *http.Request, filepath string) {
    // 设置 Content-Type
    contentType := _contentType.Get(filepath)
    w.Header().Set("Content-Type", contentType)
    
    // 读取并返回文件...
}
```

### 文件上传

```go
import "_contentType"

func UploadFile(filename string, data []byte) {
    // 获取文件类型
    contentType := _contentType.Get(filename)
    
    // 保存文件时记录类型
    file := &File{
        Name:        filename,
        ContentType: contentType,
        Data:        data,
    }
    // ...
}
```

### 静态资源服务

```go
import (
    "_contentType"
    "github.com/gin-gonic/gin"
)

func StaticHandler(c *gin.Context) {
    filename := c.Param("filename")
    
    // 自动设置正确的 Content-Type
    c.Header("Content-Type", _contentType.Get(filename))
    c.File("./static/" + filename)
}
```

---

## 🎯 典型场景

### 场景 1：静态文件服务器

```go
// 根据文件类型返回正确的 Content-Type
_contentType.Get("style.css")       // text/css
_contentType.Get("app.js")          // text/javascript
_contentType.Get("logo.png")        // image/png
_contentType.Get("font.woff2")      // font/woff2
```

### 场景 2：文件下载

```go
func DownloadHandler(w http.ResponseWriter, filename string) {
    w.Header().Set("Content-Type", _contentType.Get(filename))
    w.Header().Set("Content-Disposition", "attachment; filename="+filename)
    // 发送文件...
}
```

### 场景 3：API 响应

```go
func ExportData(format string) {
    var contentType string
    switch format {
    case "json":
        contentType = _contentType.Get("file.json")  // application/json
    case "csv":
        contentType = _contentType.Get("file.csv")   // text/csv
    case "pdf":
        contentType = _contentType.Get("file.pdf")   // application/pdf
    }
    // 设置响应头...
}
```

---

## ⚡ 性能

- ✅ **高性能**: 使用 switch case，O(1) 时间复杂度
- ✅ **无内存分配**: 返回常量字符串
- ✅ **线程安全**: 纯函数，无状态

---

## ⚠️ 注意

1. **只识别扩展名**: 不读取文件内容，仅根据扩展名判断
2. **大小写不敏感**: .JPG 和 .jpg 结果相同
3. **未知类型**: 返回 "application/octet-stream"（二进制流）

---

**License:** MIT  
**Version:** 1.0

