package handlers

import (
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"slic/internal/utils"
	"slic/internal/awss3"
	"mime/multipart"
	"strconv"
	"strings"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Invalid convert format: %s", convertFormat),
			"status":  "error",
			"error":   fmt.Sprintf("Supported formats are: %s", strings.Join(utils.GetKeys(validFormats), ", ")),
		})
	}

	s3Bucket := os.Getenv("AWS_S3_BUCKET")
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse multipart form",
			"status":  "error",
			"error":   err.Error(),
		})
	}

	files := form.File["files"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No files uploaded",
			"status":  "error",
			"error":   "No files uploaded",
		})
	}

	uploader, err := utils.AwsS3Uploader(s3Bucket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to initialize S3 uploader",
			"status":  "error",
			"error":   err.Error(),
		})
	}

	uploaded := []map[string]string{}
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to open file",
				"status":  "error",
				"error":   err.Error(),
			})
		}

		defer src.Close()

		finalFile := src
		fileName := file.Filename

		if convertFormat != "" {
			converted, _, err := utils.ConvertImage(src, convertFormat)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Failed to convert image",
					"status":  "error",
					"error":   err.Error(),
				})
			}
			finalFile = converted
			fileName = utils.ChangeExtension(fileName, convertFormat)
		}

		url, err := uploader.UploadFile(finalFile, fileName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to upload file",
				"status":  "error",
				"error":   err.Error(),
			})
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
