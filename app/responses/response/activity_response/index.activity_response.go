package activity_response

import (
	"time"
	"triesdi/app/repository/activity_repository"
)

type ActivityResponse struct {
	ActivityId         string `json:"activityId"`
	ActivityType       string `json:"activityType"`
	DoneAt             string `json:"doneAt"`
	DurationInMinutes  int    `json:"durationInMinutes"`
	CaloriesBurned     int    `json:"caloriesBurned"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}

func FormatCreateUpdateResponse (activity activity_repository.ReturnActivity) ActivityResponse {
	formatter := ActivityResponse{
		ActivityId:        activity.ID,
		ActivityType:      activity.ActivityType,
		DoneAt:            activity.DoneAt.Format(time.RFC3339Nano),
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:         activity.UpdatedAt.Format(time.RFC3339Nano),
	}
	return formatter
}

func FormatGetAllResponse (activities []activity_repository.ReturnActivity) []ActivityResponse {
	var formatter []ActivityResponse

	for _, activity := range activities {
		formatterActivity := ActivityResponse{
			ActivityId:        activity.ID,
			ActivityType:      activity.ActivityType,
			DoneAt:            activity.DoneAt.Format(time.RFC3339Nano),
			DurationInMinutes: activity.DurationInMinutes,
			CaloriesBurned:    activity.CaloriesBurned,
			CreatedAt:         activity.CreatedAt.Format(time.RFC3339Nano),
			UpdatedAt:         activity.UpdatedAt.Format(time.RFC3339Nano),
		}
		formatter = append(formatter, formatterActivity)
	}

	return formatter
}