package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Environment              string        `mapstructure:"ENVIRONMENT"`
	DBDriver                 string        `mapstructure:"DB_DRIVER"`
	DBSource                 string        `mapstructure:"DB_SOURCE"`
	MigrationURL             string        `mapstructure:"MIGRATION_URL"`
	RedisAddress             string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress        string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress        string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey        string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration      time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration     time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName          string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress       string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword      string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	AwsRegion                string        `mapstructure:"AWS_REGION"`
	AwsAccount               string        `mapstructure:"AWS_ACCOUNT"`
	S3BucketNameUserProfile  string        `mapstructure:"S3_BUCKET_NAME_USER_PROFILE"`
	S3BucketNameTransferExam string        `mapstructure:"S3_BUCKET_NAME_TRANSFER_EXAM"`
	SQSEmailSendingQueue     string        `mapstructure:"SQS_EMAIL_SENDING_QUEUE"`
	CloudFrontKeyId          string        `mapstructure:"CLOUD_FRONT_KEY_ID"`
	CloudFrontPrivateKey     string        `mapstructure:"CLOUD_FRONT_PRIVATE_KEY"`
	AWSAccessKey             string        `mapstructure:"AWS_ACCESS_KEY"`
	AWSSecretAccessKey       string        `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	FrontEndSubDomain        string        `mapstructure:"FRONTEND_SUB_DOMAIN"`
	FrontEndDomain           string        `mapstructure:"FRONTEND_DOMAIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
