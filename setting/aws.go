package setting

type AWsSetting struct {
	Connect         string
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Bucket          string
}

var Aws = &AWsSetting{}

func AwsSetup() *AWsSetting {
	Aws.Connect = env("AWS_CONNECT", "false")
	Aws.AccessKeyID = env("AWS_ACCESS_KEY_ID", "")
	Aws.SecretAccessKey = env("AWS_SECRET_ACCESS_KEY", "")
	Aws.Region = env("AWS_REGION", "")
	Aws.Bucket = env("AWS_BUCKET", "")
	return Aws
}
