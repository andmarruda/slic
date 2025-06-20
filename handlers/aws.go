package handlers

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"slic/internal/awss3"
	"slic/internal/errors"
	"slic/internal/utils"
)

func Upload(c *fiber.Ctx) error {
	convertFormat := strings.ToLower(c.FormValue("convertFormat", ""))

	validFormats := map[string]bool{
		"jpg":  true,
		"jpeg": true,
		"png":  true,
		"gif":  true,
		"webp": true,
	}

	if convertFormat != "" && !validFormats[convertFormat] {
		return errors.ErrorJSON(c, fiber.StatusBadRequest, "Invalid convertFormat",
			fmt.Errorf("Supported formats are: %s", strings.Join(utils.MapKeys(validFormats), ", ")))
	}

	s3Bucket := os.Getenv("AWS_S3_BUCKET")
	form, err := c.MultipartForm()
	if err != nil {
		return errors.ErrorJSON(c, fiber.StatusBadRequest, "Failed to parse multipart form", err)
	}

	files := form.File["files"]
	if len(files) == 0 {
		return errors.ErrorJSON(c, fiber.StatusBadRequest, "No files uploaded", fmt.Errorf("Please upload at least one file"))
	}

	uploader, err := awss3.S3Uploader(s3Bucket)
	if err != nil {
		return errors.ErrorJSON(c, fiber.StatusInternalServerError, "Failed to initialize S3 uploader", err)
	}

	uploaded := []map[string]string{}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return errors.ErrorJSON(c, fiber.StatusInternalServerError, "Failed to open file", err)
		}
		defer src.Close()

		finalFile := src
		fileName := file.Filename

		if convertFormat != "" {
			converted, _, err := utils.ConvertImage(src, convertFormat)
			if err != nil {
				return errors.ErrorJSON(c, fiber.StatusInternalServerError, "Failed to convert image", err)
			}
			finalFile = converted
			fileName = utils.ChangeExtension(fileName, convertFormat)
		}

		contentType := utils.GetContentTypeFromFilename(fileName)

		url, err := uploader.Upload(c.Context(), fileName, finalFile, contentType)
		if err != nil {
			return errors.ErrorJSON(c, fiber.StatusInternalServerError, "Failed to upload file", err)
		}

		uploaded = append(uploaded, map[string]string{
			"fileName": fileName,
			"url":      url,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Files uploaded successfully",
		"status":  "success",
		"data":    uploaded,
	})
}
