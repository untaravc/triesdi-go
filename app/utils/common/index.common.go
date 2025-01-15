package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"triesdi/app/cache"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashingPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash), err
}

func ConsoleLog(data ...interface{}) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Print the formatted JSON
	fmt.Println(string(jsonBytes))
}

// GetCaloriesPerMinute retrieves the calories per minute for a valid activity type
func GetCaloriesPerMinute(activityType string) (int, error) {
	calories, exists := cache.ActivityTypeCache[activityType]
	if !exists {
		return 0, errors.New("invalid activity type")
	}
	return calories, nil
}