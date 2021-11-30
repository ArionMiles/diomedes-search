package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func SendToQueue(sqsClient *sqs.Client, queueURL, msg string) (*sqs.SendMessageOutput, error) {
	resp, err := sqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
		MessageBody: aws.String(msg),
		QueueUrl:    aws.String(queueURL),
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
