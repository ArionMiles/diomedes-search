package models

import "github.com/ArionMiles/gobms"

type Reminder struct {
	Id          string
	ChatID      int64
	Completed   bool
	Date        string
	Format      string
	Language    string
	MovieName   string
	RegionCode  string
	RegionName  string
	TheaterCode string
}

type Result struct {
	Reminder Reminder
	Shows    []gobms.Show
}
