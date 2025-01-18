package activity_repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	UpdateActivity(id string, activity Activity) (ReturnActivity, error)
	GetActivity(id string, filters ActivityFilter) ([]ReturnActivity, error)
	DeleteActivity(id string) (Activity, error)
	CreateActivity(activity Activity) (ReturnActivity, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateActivity(id string, activity Activity) (ReturnActivity, error) {
	// Perform the update operation
	if err := r.db.Model(&Activity{}).Where("id = ?", id).Updates(activity).Error; err != nil {
		return ReturnActivity{}, err
	}

	// Retrieve the updated data
	var updatedActivity ReturnActivity
	if err := r.db.Model(&Activity{}).Where("id = ?", id).First(&updatedActivity).Error; err != nil {
		return ReturnActivity{}, err
	}

	return updatedActivity, nil
}


func (r *repository) GetActivity(user_id string, filters ActivityFilter) ([]ReturnActivity, error) {
	query := r.db.Model(&Activity{})

	// Apply filters
	if filters.ActivityType != "" {
		query = query.Where("activity_type = ?", filters.ActivityType)
	}

	if filters.DoneAtFrom != "" {
		query = query.Where("done_at >= ?", filters.DoneAtFrom)
		fmt.Printf("querys from: %v\n", "done_at >= "+filters.DoneAtFrom)
	}

	if filters.DoneAtTo != "" {
		query = query.Where("done_at <= ?", filters.DoneAtTo)
		fmt.Printf("querys to: %v\n", "done_at <= "+filters.DoneAtTo)
	}

	if filters.CaloriesBurnedMin != nil {
		query = query.Where("calories_burned >= ?", filters.CaloriesBurnedMin)
	}

	if filters.CaloriesBurnedMax != nil {
		query = query.Where("calories_burned <= ?", filters.CaloriesBurnedMax)
	}

	// Filter by user_id
	query = query.Where("user_id = ?", user_id)

	// Pagination
	query = query.Limit(filters.Limit).Offset(filters.Offset)

	var activities []ReturnActivity
	if err := query.Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}

func (r *repository) CreateActivity(activity Activity) (ReturnActivity, error) {
    if err := r.db.Create(&activity).Error; err != nil {
        return ReturnActivity{}, err
    }

    var createdActivity ReturnActivity
	if err := r.db.Table("activities").First(&createdActivity, "id = ?", activity.ID).Error; err != nil {
        return ReturnActivity{}, err
    }

    return createdActivity, nil
}

func (r *repository) DeleteActivity(id string) (Activity, error) {
	var activity Activity
	if err := r.db.Where("id = ?", id).First(&activity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return activity, errors.New("activity not found")
		}
		return activity, err
	}
	if err := r.db.Delete(&activity).Error; err != nil {
		return activity, err
	}
	return activity, nil
}
