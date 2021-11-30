package notification

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"gopkg.in/tucnak/telebot.v2"

	"github.com/ArionMiles/diomedes-search/internal/models"
	"github.com/ArionMiles/diomedes-search/internal/utils"
	"github.com/ArionMiles/gobms"
)

func formatShowtimes(result models.Result) string {
	formattedShows := make([]string, len(result.Shows))
	for index, item := range result.Shows {
		showURL, err := gobms.GetShowtimeURL(result.Reminder.TheaterCode, item.SessionID)
		if err != nil {
			log.Printf("Cannot get Showtime URL for %v", result.Reminder.MovieName)
		}
		date, err := utils.Iso8601ToDate(result.Reminder.Date)
		formattedDate := date.Format("02 January 2006")
		if err != nil {
			log.Printf("Incorrect date format: %s", result.Reminder.Date)
		}
		formattedShows[index] = fmt.Sprintf("%v (%v) \n [%v](%v)",
			result.Reminder.MovieName,
			formattedDate,
			item.ShowTimeDisplay,
			showURL,
		)
	}
	return strings.Join(formattedShows, "\n")
}

func SendToTelegram(client *telebot.Bot, result models.Result) (*telebot.Message, error) {
	messageText := formatShowtimes(result)
	msg, err := client.Send(
		telebot.ChatID(result.Reminder.ChatID),
		messageText,
		telebot.ModeMarkdown,
		telebot.NoPreview,
	)
	if err != nil {
		return nil, err
	}
	// Pin message to notify everyone in a group
	// Only works for groups
	err = client.Pin(msg)
	if err != nil {
		log.Println("Failed to pin message:", err)
	}
	return msg, nil
}

func MarkDone(client *dynamodb.Client, tableName, reminderID string) (*dynamodb.UpdateItemOutput, error) {
	result, err := client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: reminderID},
		},
		UpdateExpression: aws.String("set Completed = :completed"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":completed": &types.AttributeValueMemberBOOL{Value: true},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
