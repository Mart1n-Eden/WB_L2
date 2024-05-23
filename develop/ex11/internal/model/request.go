package model

import (
	"time"
)

type Request struct {
	UserId    int    `json:"user_id"`
	EventName string `json:"event_name"`
	Date      string `json:"date"`
}

func CastToEvent(request Request) (Event, error) {
	date, err := time.Parse("2006-01-02 15:04:05", request.Date)

	if err != nil {
		return Event{}, err
	}

	return Event{Name: request.EventName, Date: date}, nil
}
