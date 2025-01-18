package activity_repository

import (
	"time"
)

type Activity struct {
	ID string
	UserID string
	ActivityType string
	DoneAt string
	DurationInMinutes int
	CaloriesBurned int
}

type ReturnActivity struct {
	ID string
	UserID string
	ActivityType string
	DoneAt time.Time
	DurationInMinutes int
	CaloriesBurned int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ActivityFilter struct {
	Limit             int
	Offset            int
	ActivityType      string
	DoneAtFrom        string
	DoneAtTo          string
	CaloriesBurnedMin *int
	CaloriesBurnedMax *int
}