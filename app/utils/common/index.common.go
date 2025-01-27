package common

import (
	"encoding/json"
	"fmt"
	"math/rand"

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

var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomString(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
