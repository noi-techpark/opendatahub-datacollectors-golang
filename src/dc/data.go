package dc

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func GetData() []byte {
	slog.Info("Getting file from S3 bucket...")

	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET_NAME")
	file := os.Getenv("AWS_FILE_NAME")

	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")

	buf := aws.NewWriteAtBuffer([]byte{})
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		slog.Error("session error", err)
	}

	downloader := s3manager.NewDownloader(session)
	numBytes, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(file),
		})
	if err != nil {
		slog.Error("download error", err)
	}

	slog.Info("Getting file done. " + strconv.FormatInt(numBytes, 10) + " bytes counted.")

	return buf.Bytes()
}
