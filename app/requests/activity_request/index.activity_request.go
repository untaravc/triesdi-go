package activity_request

type ActivityRequest struct {
	ActivityType      string `json:"activityType" validate:"required"`
	DoneAt            string `json:"doneAt" validate:"required"`
	DurationInMinutes int    `json:"durationInMinute" validate:"required"`
}
