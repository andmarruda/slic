package utils

import (
	"strings"
	"mime"
	"path/filepath"
)

func changeExtension(filename, newExt string) string {
    if dot := strings.LastIndex(filename, "."); dot != -1 {
        return filename[:dot] + "." + newExt
    }
    return filename + "." + newExt
}

func GetContentTypeFromFilename(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return "application/octet-stream"
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream"
	}

	return mimeType
}
