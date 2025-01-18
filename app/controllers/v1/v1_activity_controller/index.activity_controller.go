package v1_activity_controller

import (
	"net/http"
	"triesdi/app/configs/db_config"
	"triesdi/app/repository/activity_repository"
	"triesdi/app/requests/activity_request"
	"triesdi/app/responses/response/activity_response"
	"triesdi/app/service/activity_service"
	"triesdi/app/validator"

	"github.com/gin-gonic/gin"
)

func GetActivities(c *gin.Context) {

	db := db_config.GetDB()
	activityRepository := activity_repository.NewRepository(db)
	activityService := activity_service.NewService(activityRepository)

	activities, err := activityService.GetActivity(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Formatting activities to Response
	activitiesResponse := activity_response.FormatGetAllResponse(activities)

	c.JSON(http.StatusOK, activitiesResponse)
}

func CreateActivity(ctx *gin.Context) {
	user_id, exist := ctx.Get("id")
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	activityRequest := new(activity_request.ActivityRequest)
	if err := ctx.ShouldBindJSON(activityRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validator.ValidateStruct(activityRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// VALIDATE Activity Type
	if err := validator.ValidateActivityType(activityRequest.ActivityType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Activity Type"})
		return
	}

	db := db_config.GetDB()
	activityRepository := activity_repository.NewRepository(db)
	activityService := activity_service.NewService(activityRepository)

	var activity activity_repository.ReturnActivity
	var activityResponse activity_response.ActivityResponse
	var err error

	activity, err = activityService.CreateActivity(user_id.(string), *activityRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	activityResponse = activity_response.FormatCreateUpdateResponse(activity)

	ctx.JSON(http.StatusCreated, activityResponse)
}

func UpdateActivity(ctx *gin.Context) {
	
	user_id, exist := ctx.Get("id")
	if !exist {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	data_id := ctx.Param("id")
	if data_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	activityRequest := new(activity_request.ActivityRequest)
	if err := ctx.ShouldBindJSON(activityRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := validator.ValidateStruct(activityRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// VALIDATE Activity Type
	if err := validator.ValidateActivityType(activityRequest.ActivityType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Activity Type"})
		return
	}

	db := db_config.GetDB()
	activityRepository := activity_repository.NewRepository(db)
	activityService := activity_service.NewService(activityRepository)

	var activity activity_repository.ReturnActivity
	var activityResponse activity_response.ActivityResponse
	var err error

	activity, err = activityService.UpdateActivity(user_id.(string), data_id,*activityRequest)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	activityResponse = activity_response.FormatCreateUpdateResponse(activity)

	ctx.JSON(http.StatusOK, activityResponse)
}

func DeleteActivity(ctx *gin.Context) {
	data_id := ctx.Param("id")
	if data_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	db := db_config.GetDB()
	activityRepository := activity_repository.NewRepository(db)
	activityService := activity_service.NewService(activityRepository)

	var err error

	_, err = activityService.DeleteActivity(data_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
