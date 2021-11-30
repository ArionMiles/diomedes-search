package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ArionMiles/diomedes-search/internal/models"
	"github.com/ArionMiles/diomedes-search/internal/notification"
	"github.com/ArionMiles/diomedes-search/internal/utils"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) {
	bot, err := utils.GetTGClient()
	if err != nil {
		log.Panic("Unable to instantiate telegram client", err)
	}
	ddbClient := utils.GetDDBClient()
	ddbTable := os.Getenv("DDB_TABLE")

	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s", message.MessageId, message.EventSource)
		result := models.Result{}
		err := json.Unmarshal([]byte(message.Body), &result)
		if err != nil {
			log.Print("Unable to unmarshall message", err)
			continue
		}

		msg, err := notification.SendToTelegram(bot, result)
		if err != nil {
			log.Print("Sending message failed", err)
			log.Print("ChatID: ", result.Reminder.ChatID)
			continue
		}
		log.Print("Bot message ID:", msg.ID)

		updateOutput, err := notification.MarkDone(ddbClient, ddbTable, result.Reminder.Id)
		if err != nil {
			log.Printf("Error updating %v. Error: %v", result.Reminder.Id, err)
			continue
		}
		log.Print(updateOutput)
	}

}

func main() {
	lambda.Start(handler)
}
