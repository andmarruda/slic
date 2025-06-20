package awss3

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	Client *s3.Client
	Bucket string
}

// Inicializa o uploader com client e nome do bucket
func NewS3Uploader(bucket string) (*S3Uploader, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	return &S3Uploader{
		Client: client,
		Bucket: bucket,
	}, nil
}

// Upload realiza o upload e retorna a URL final do arquivo
func (u *S3Uploader) Upload(ctx context.Context, filename string, body io.Reader, contentType string) (string, error) {
	key := fmt.Sprintf("uploads/%d-%s", time.Now().UnixNano(), filename)

	_, err := u.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.Bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
		ACL:         aws.String("public-read"), // opcional: garante visibilidade pública
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.Bucket, key)
	return url, nil
}
