package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/ArionMiles/diomedes-search/internal/models"
	"github.com/ArionMiles/diomedes-search/internal/queue"
	"github.com/ArionMiles/diomedes-search/internal/search"
	"github.com/ArionMiles/diomedes-search/internal/utils"
)

func handler(ctx context.Context, b json.RawMessage) {
	ddb := utils.GetDDBClient()
	sqsClient := utils.GetSQSClient()
	queueURL := os.Getenv("QUEUE_URL")
	ddbTable := os.Getenv("DDB_TABLE")

	result, err := ddb.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:        aws.String(ddbTable),
		FilterExpression: aws.String("Completed = :completed"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":completed": &types.AttributeValueMemberBOOL{Value: false},
		},
	})
	if err != nil {
		panic(err)
	}

	reminders := []models.Reminder{}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &reminders)
	if err != nil {
		log.Fatal("Unable to unmarshal reminders", err)
	}

	for _, item := range reminders {
		log.Println(item.MovieName, item.Completed)
	}

	for _, item := range reminders {
		shows, err := search.FindShows(item)
		if err != nil {
			log.Printf("No shows found for %s", item.MovieName)
			continue
		}
		showsJSON, _ := json.MarshalIndent(shows, "", "\t")
		message := string(showsJSON)
		msgOutput, err := queue.SendToQueue(sqsClient, queueURL, message)
		if err != nil {
			log.Println("Got an error sending the message:")
			log.Println(err)
			continue
		}
		log.Println("Sent message with ID:", *msgOutput.MessageId)
	}
}

func main() {
	lambda.Start(handler)
}
