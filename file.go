package utilities

import (
	"mime"
	"path/filepath"
	"strings"
)

// MimeTypeFromExtension returns the MIME type based on the file extension.
// Defaults to "application/octet-stream" if unknown.
func MimeTypeFromExtension(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	if ext == "" {
		return "application/octet-stream"
	}
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream"
	}
	return mimeType
}
