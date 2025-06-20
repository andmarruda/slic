# SLIC Image Upload & Conversion API

This microservice provides a serverless-ready HTTP API to:

- Upload one or more images via `multipart/form-data`
- Optionally convert each image to another format (e.g., JPG → WebP)
- Store the resulting images in AWS S3
- Return public URLs of the uploaded images

Built with **Go**, **Fiber**, and **AWS SDK v2**, it's optimized for fast, flexible image handling without persistent storage.

---

## 📦 Features

- ✅ Upload multiple images in a single request
- 🔄 Optional real-time format conversion (JPG, PNG, WebP, GIF, JPEG)
- ☁️ Uploads to AWS S3 with `public-read` access
- ⚡ Lightweight, serverless-compatible
- 🔐 Clean error handling and structured responses

---

## 🚀 API Endpoint

### `POST /api/v1/aws-s3/upload`

Uploads one or more images. Optionally converts them to the format specified via `convertFormat`.

#### 📤 Request

- **Content-Type**: `multipart/form-data`
- **Fields**:
  - `files[]`: one or more image files
  - `convertFormat` *(optional)*: one of `jpg`, `jpeg`, `png`, `webp`, `gif`

#### 📥 Example cURL

```bash
curl -X POST http://localhost:8080/api/v1/aws-s3/upload \
  -F "files[]=@photo1.jpg" \
  -F "files[]=@photo2.png" \
  -F "convertFormat=webp"

## Response

### 🟢 Success (HTTP 200)

```
{
  "message": "Files uploaded successfully",
  "status": "success",
  "data": [
    {
      "fileName": "photo1.webp",
      "url": "https://your-bucket.s3.amazonaws.com/uploads/1689487000-photo1.webp"
    },
    {
      "fileName": "photo2.webp",
      "url": "https://your-bucket.s3.amazonaws.com/uploads/1689487000-photo2.webp"
    }
  ]
}
```

### 🔴 Error (HTTP 4xx / 5xx)

```
{
  "message": "Invalid convertFormat",
  "status": "error",
  "error": "Supported formats are: jpg, jpeg, png, gif, webp"
}
```

## 🛠️ Environment Variables

| Variable | | Description |
| -------- | | ----------- |
| AWS_ACCESS_KEY_ID | Your AWS access key |
| AWS_SECRET_ACCESS_KEY | Your AWS secret |
| AWS_REGION | AWS region (e.g. us-east-1) |
| AWS_S3_BUCKET | Target S3 bucket name |

## 🧠 How It Works

Start the server (assuming main.go runs your app):

```
go run main.go

```

Then test using tools like:

- curl
- Postman
- Your frontend app

## 📁 Project Structure

slic/
├── internal/
│   ├── awss3/        # S3 uploader logic
│   ├── convert/      # Image conversion logic
│   └── utils.go      # Helpers (content-type, extension)
├── handlers/         # HTTP handler logic
│   └── upload.go
├── routes/           # Route groupings
│   └── routes.go
├── main.go           # App entry point
├── go.mod / go.sum   # Go module dependencies

## 🧑‍💻 License

MIT License © 2024 Anderson Arruda

# 🖼️ SLIC Image Upload & Conversion API

[![Go Version](https://img.shields.io/badge/Go-1.22-blue)](https://go.dev/doc/go1.22)
[![Fiber Framework](https://img.shields.io/badge/Fiber-2.x-green)](https://gofiber.io/)
[![License](https://img.shields.io/badge/license-MIT-lightgrey)](LICENSE)
[![Docker Ready](https://img.shields.io/badge/docker-ready-blue)](https://www.docker.com/)

---

A high-performance microservice to:

- 🗂️ Upload one or more images
- 🔄 Optionally convert them to another format (JPG, PNG, WebP, etc.)
- ☁️ Store in AWS S3
- 📎 Return public URLs

Built with **Go**, **Fiber**, and **AWS SDK v2** – stateless and serverless-ready.

---

<details>
<summary><strong>📦 Features</strong></summary>

- ✅ Upload multiple files via `multipart/form-data`
- 🔄 Optional image format conversion (WebP, PNG, JPG, GIF, JPEG)
- ☁️ Upload to AWS S3 with public-read access
- ⚡ Lightweight and production-ready
- 🔐 Structured error responses
- 🐳 Docker-compatible

</details>

---

<details>
<summary><strong>🚀 API Endpoint</strong></summary>

### `POST /api/v1/aws-s3/upload`

Uploads images to S3, with optional format conversion.

#### 🔧 Request

- **Content-Type**: `multipart/form-data`
- **Fields**:
  - `files[]`: one or more image files
  - `convertFormat` *(optional)*: `jpg`, `jpeg`, `png`, `webp`, or `gif`

#### 📥 Example cURL

```bash
curl -X POST http://localhost:8080/api/v1/aws-s3/upload \
  -F "files[]=@image1.jpg" \
  -F "files[]=@image2.png" \
  -F "convertFormat=webp"
