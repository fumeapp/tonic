package aws

import (
	"bytes"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"

	"github.com/fumeapp/tonic/setting"
)

func cfg() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return cfg, errors.New("failed to load AWS config " + err.Error())
	}

	return cfg, nil
}

func S3() *s3.Client {
	cfg, err := cfg()
	if err != nil {
		log.Fatalf("failed to load AWS config, %v", err)
	}
	return s3.NewFromConfig(cfg)
}

func Uploader() *manager.Uploader {
	return manager.NewUploader(S3())
}

func SES() *ses.Client {
	cfg, err := cfg()
	if err != nil {
		log.Fatalf("failed to load AWS config, %v", err)
	}
	return ses.NewFromConfig(cfg)
}

func randToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("failed to generate random token, %v", err)
	}
	return fmt.Sprintf("%x", b)
}

// Uploads a file to S3 naming it after a hash of the file contents.
// Accepts a public URL
// returns the URL of the uploaded file and an error if there was one.
func Upload(url string) (string, error) {

	var extension string

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(bodyBytes)

	switch contentType {
	case "image/jpeg":
		extension = "jpg"
	case "image/png":
		extension = "png"
	default:
		return "", errors.New("unable to detect Content Type: " + contentType)
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
