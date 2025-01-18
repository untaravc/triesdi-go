package activity_response

import (
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
		DoneAt:            activity.DoneAt.Format("2006-01-02T15:04:05.000Z07:00")[:23] + "Z",
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")[:23] + "Z",
		UpdatedAt:         activity.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")[:23] + "Z",
	}
	return formatter
}

func FormatGetAllResponse (activities []activity_repository.ReturnActivity) []ActivityResponse {
	var formatter []ActivityResponse

	for _, activity := range activities {
		formatterActivity := ActivityResponse{
			ActivityId:        activity.ID,
			ActivityType:      activity.ActivityType,
			DoneAt:            activity.DoneAt.Format("2006-01-02T15:04:05.000Z07:00")[:23] + "Z",
			DurationInMinutes: activity.DurationInMinutes,
			CaloriesBurned:    activity.CaloriesBurned,
			CreatedAt:         activity.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")[:23] + "Z",
			UpdatedAt:         activity.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")[:23] + "Z",
		}
		formatter = append(formatter, formatterActivity)
	}

	return formatter
}