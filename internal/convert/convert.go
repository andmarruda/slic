package convert

import {
	"bytes"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"strings"

	"github.com/disintegration/imaging"
}

func convertImage(file multipart.File, format string) (io.Reader, string, erro) {
	img, formatDetected, err := image.Decode(file)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode image: %w", err)
	}

	_ = formatDetected
	var buf bytes.Buffer

	format = strings.ToLower(format)

	switch format {
		case "jpg", "jpeg":
			err = imaging.Encode(&buf, img, imaging.JPEG)
		case "png":
			err = imaging.Encode(&buf, img, imaging.PNG)
		case "gif":
			err = imaging.Encode(&buf, img, imaging.GIF)
		case "webp":
			err = imaging.Encode(&buf, img, imaging.WebP)
		default:
			return nil, "", fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return nil, "", fmt.Errorf("failed to encode image: %w", err)
	}

	return &buf, format, nil
}