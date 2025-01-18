package activity_service

import (
	"strconv"
	"triesdi/app/cache"
	"triesdi/app/repository/activity_repository"
	"triesdi/app/requests/activity_request"
	"triesdi/app/responses/response/activity_response"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
    UpdateActivity(user_id string, activity activity_request.ActivityRequest) (activity_response.ActivityResponse, error)
    GetActivity(c *gin.Context) ([]activity_response.ActivityResponse, error)
    CreateActivity(activity activity_request.ActivityRequest) (activity_response.ActivityResponse, error)
	DeleteActivity(user_id string) (activity_response.ActivityResponse, error)
}

type service struct {
    repo activity_repository.Repository
}

func NewService(repo activity_repository.Repository) *service {
    return &service{repo}
}

func (s *service) UpdateActivity(user_id string, id string, activity activity_request.ActivityRequest) (activity_repository.ReturnActivity, error) {
	_, err := time.Parse(time.RFC3339, activity.DoneAt)
	if err != nil {
		return activity_repository.ReturnActivity{}, err
	}

    activityToUpdate := activity_repository.Activity{
        UserID:            user_id,
        ActivityType:      activity.ActivityType,
        DoneAt:            activity.DoneAt,
        DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    calculateCaloreisBurned(activity.DurationInMinutes, activity.ActivityType),
    }

    updatedActivity, err := s.repo.UpdateActivity(id, activityToUpdate)
    if err != nil {
        return activity_repository.ReturnActivity{}, err
    }

    return updatedActivity, nil
}

func (s *service) GetActivity(c *gin.Context) ([]activity_repository.ReturnActivity, error) {
    // Parse query parameters with defaults
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil || limit <= 0 {
		limit = 5
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}

	activityType := c.Query("activityType") // No need to validate, invalid value will be ignored

	// Parse and validate doneAtFrom
	doneAtFromStr := c.Query("doneAtFrom")
	// fmt.Printf("doneAtFromStr: %v\n", doneAtFromStr)
	// if doneAtFromStr != "" {
	// 	_, err := time.Parse(time.RFC3339Nano, doneAtFromStr)
	// 	if err != nil {
	// 		doneAtFromStr = ""
	// 	}
	// }

	// Parse and validate doneAtTo
	doneAtToStr := c.Query("doneAtTo")
	// fmt.Printf("doneAtToStr: %v\n", doneAtToStr)
	// if doneAtToStr != "" {
	// 	if _, err := time.Parse(time.RFC3339Nano, doneAtToStr); err != nil {
	// 		doneAtToStr = ""
	// 	}
	// }

	// Parse and validate caloriesBurnedMin
	var caloriesBurnedMin *int
	caloriesBurnedMinStr := c.Query("caloriesBurnedMin")
	if caloriesBurnedMinStr != "" {
		value, err := strconv.Atoi(caloriesBurnedMinStr)
		if err == nil && value >= 0 {
			caloriesBurnedMin = &value
		}
	}

	// Parse and validate caloriesBurnedMax
	var caloriesBurnedMax *int
	caloriesBurnedMaxStr := c.Query("caloriesBurnedMax")
	if caloriesBurnedMaxStr != "" {
		value, err := strconv.Atoi(caloriesBurnedMaxStr)
		if err == nil && value >= 0 {
			caloriesBurnedMax = &value
		}
	}
	
	// Construct filter object
	filters := activity_repository.ActivityFilter{
		Limit:              limit,
		Offset:             offset,
		ActivityType:       activityType,
		DoneAtFrom:         doneAtFromStr,
		DoneAtTo:           doneAtToStr,
		CaloriesBurnedMin:  caloriesBurnedMin,
		CaloriesBurnedMax:  caloriesBurnedMax,
	}
	
	// Get user_id from context
	userID, exist := c.Get("id")
	if !exist {
		return []activity_repository.ReturnActivity{}, nil
	}
	user_id, ok := userID.(string)
	if !ok {
		return []activity_repository.ReturnActivity{}, nil
	}

	activities, err := s.repo.GetActivity(user_id, filters)
	if err != nil {
		return []activity_repository.ReturnActivity{}, err
	}

	return activities, nil
}

func (s *service) CreateActivity(user_id string, activity activity_request.ActivityRequest) (activity_repository.ReturnActivity, error) {

	_, err := time.Parse(time.RFC3339, activity.DoneAt)
	if err != nil {
		return activity_repository.ReturnActivity{}, err
	}

    newActivity := activity_repository.Activity{
		ID: uuid.New().String(),
        UserID:            user_id,
        ActivityType:      activity.ActivityType,
        DoneAt:            activity.DoneAt,
        DurationInMinutes: activity.DurationInMinutes,
        CaloriesBurned:    calculateCaloreisBurned(activity.DurationInMinutes, activity.ActivityType),
    }

    createdActivity, err := s.repo.CreateActivity(newActivity)
    if err != nil {
        return activity_repository.ReturnActivity{}, err
    }

	return createdActivity, nil
}

func (s *service) DeleteActivity(id string) (activity_repository.Activity, error) {
	activity, err := s.repo.DeleteActivity(id)
	if err != nil {
		return activity_repository.Activity{}, err
	}

	return activity, nil
}

func calculateCaloreisBurned(durationInMinutes int, activity_type string) int {
	cache.GetCaloriesPerMinute(activity_type)
	return durationInMinutes * cache.GetCaloriesPerMinute(activity_type)
}