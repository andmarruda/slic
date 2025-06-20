# SLIC - Serverless Lightweight Image Converter

**SLIC** is a serverless microservice for image format conversion and CDN upload. It receives one or more images via HTTP, converts each to the desired format (e.g., WebP, PNG, JPEG), uploads them to a configured CDN, and returns the final public URLs.

## Features

- 🔄 Convert one or multiple images on the fly
- 🎯 Supports WebP, JPEG, PNG, and other common formats
- ☁️ Uploads to your CDN (e.g., AWS S3)
- ⚡️ Serverless and stateless – no database required
- 🔐 Optional authentication token for secure usage

## Use Case

- Automatically convert user-uploaded images to optimized formats
- Resize and reformat before delivery to frontend apps
- Decouple image handling from your main application

## API

### Endpoint

```http
POST /convert
