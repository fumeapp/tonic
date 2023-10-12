package aws

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"

	"github.com/fumeapp/tonic/setting"
)

func Config(params ...string) aws.Config {

	region := "us-east-1"

	// check for a starting parameter of an aws region
	if len(params) > 0 {
		if params[0] != "" {
			region = params[0]
		}
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("failed to load AWS config, %v", err)
	}

	return cfg
}

func S3() *s3.Client {
	return s3.NewFromConfig(Config())
}

func Uploader() *manager.Uploader {
	return manager.NewUploader(S3())
}

func SES() *ses.Client {
	return ses.NewFromConfig(Config())
}

func randToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("failed to generate random token, %v", err)
	}
	return fmt.Sprintf("%x", b)
}

// UploadURL
// Uploads a file to S3 naming it after a hash of the file contents.
// Accepts a public URL
// returns the URL of the uploaded file and an error if there was one.
func UploadURL(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	return Upload(bodyBytes)
}

// UploadFile - same functionality as UploadURL but take in a multipart.FileHeader
func UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	bodyBytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return Upload(bodyBytes)
}

// Upload
// Uploads a file to S3 naming it after a hash of the file contents.
// Accepts a public URL
// returns the URL of the uploaded file and an error if there was one.
func Upload(bodyBytes []byte) (string, error) {

	extension, contentType, err := getExtension(bodyBytes)
	if err != nil {
		return "", err
	}

	result, err := Uploader().Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(setting.Aws.Bucket),
		Key:         aws.String(randToken() + "." + extension),
		Body:        io.NopCloser(bytes.NewBuffer(bodyBytes)),
		ACL:         "public-read",
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil
}

// Figure out file extension and content type
func getExtension(bytes []byte) (string, string, error) {

	var extension string
	contentType := http.DetectContentType(bytes)

	switch contentType {
	case "image/jpg":
		extension = "jpg"
	case "image/jpeg":
		extension = "jpg"
	case "image/png":
		extension = "png"
	case "image/gif":
		extension = "gif"
	case "image/webp":
		extension = "webp"
	case "image/svg":
		extension = "svg"
	default:
		return "", "", errors.New("unable to detect Content Type: " + contentType)
	}

	return extension, contentType, nil
}

func SendEmail(to string, subject string, body string, from string) (*ses.SendEmailOutput, error) {
	return SES().SendEmail(context.TODO(), &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{
				to,
			},
		},
		Source: aws.String(from),
		Message: &types.Message{
			Subject: &types.Content{
				Data: &subject,
			},
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
		},
	})
}
