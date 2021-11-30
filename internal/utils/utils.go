package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"gopkg.in/tucnak/telebot.v2"
)

func Iso8601ToDate(isoDate string) (*time.Time, error) {
	formattedDate, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return nil, err
	}
	return &formattedDate, nil
}

func GetDDBClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	ddb := dynamodb.NewFromConfig(cfg)
	return ddb
}

func GetSQSClient() *sqs.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	sqs := sqs.NewFromConfig(cfg)
	return sqs
}

func GetTGClient() (*telebot.Bot, error) {
	bot, err := telebot.NewBot(telebot.Settings{
		Token: os.Getenv("TG_TOKEN"),
	})
	if err != nil {
		return nil, err
	}
	return bot, nil
}
