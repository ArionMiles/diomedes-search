package search

import (
	"log"

	"github.com/ArionMiles/gobms"

	"github.com/ArionMiles/diomedes-search/internal/models"
	"github.com/ArionMiles/diomedes-search/internal/utils"
)

func FindShows(r models.Reminder) (*models.Result, error) {
	bms, err := gobms.NewClient(r.RegionCode, r.RegionName)
	if err != nil {
		log.Printf("Error generating client for Code: %s and Name: %s", r.RegionCode, r.RegionName)
		return nil, err
	}
	eventCode, err := bms.GetEventCode(r.MovieName, r.Language, r.Format)
	if err != nil {
		log.Printf("No Event Code found for %s", r.MovieName)
		return nil, err
	}
	formattedDate, err := utils.Iso8601ToDate(r.Date)
	if err != nil {
		log.Printf("Incorrect date format: %s", r.Date)
		return nil, err
	}
	shows, err := bms.GetShowtimes(eventCode, r.TheaterCode, *formattedDate)
	if err != nil {
		log.Printf("No Showtimes found for %s", r.MovieName)
		return nil, err
	}

	return &models.Result{
		Reminder: r,
		Shows:    shows,
	}, nil
}
