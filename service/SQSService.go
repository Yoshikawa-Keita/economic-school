package service

import (
	"context"
	"economic-school/util"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
)

type EmailConfirmationMessage struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SendMessageToSQS(env util.Config, queueName string, message map[string]string) error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	sqsClient := sqs.NewFromConfig(cfg)

	// Encode the message to JSON
	messageJson, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	queueUrl := fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", env.AwsRegion, env.AwsAccount, queueName)
	input := &sqs.SendMessageInput{
		MessageBody: aws.String(string(messageJson)),
		QueueUrl:    aws.String(queueUrl),
	}

	_, err = sqsClient.SendMessage(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to send message to SQS: %w", err)
	}

	return nil

}
