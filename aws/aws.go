package aws

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net/http"

	_ "github.com/aws/aws-sdk-go-v2"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/fumeapp/tonic/setting"
)

var Config *config.Config
var S3 *s3.Client
var Uploader *manager.Uploader

func Setup() {
	if setting.Aws.Connect != "true" {
		return
	}
	Config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load AWS config, %v", err)
	}

	S3 = s3.NewFromConfig(Config)
	Uploader = manager.NewUploader(S3)
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

	hasher := md5.New()

	if _, err := io.Copy(hasher, response.Body); err != nil {
		return "", err
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	uploader := manager.NewUploader(S3)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(setting.Aws.Bucket),
		Key:         aws.String(hash + "." + extension),
		Body:        io.NopCloser(bytes.NewBuffer(bodyBytes)),
		ACL:         "public-read",
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return "", err
	}

	return result.Location, nil

}
