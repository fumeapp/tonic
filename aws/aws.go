package aws

import (
	"context"
	"log"

	_ "github.com/aws/aws-sdk-go-v2"
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

func UNUSED(x ...any) {}
