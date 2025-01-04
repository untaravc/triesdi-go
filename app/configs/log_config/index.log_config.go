package log_config

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var defaultLogPath = "./logs/gin.log"
var (
	currentLogFile string
	logger         *log.Logger
	mutex          sync.Mutex
)

func createLogFolderIfNotExist(path string) {
	dir := filepath.Dir(path)

	log.Println("Creating directory")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0644)
		if err != nil {
			log.Println("Fail to create directory")
		} else {
			log.Println("Directory created")
		}
	}

}

func openOrCreateLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		var errorCreateFile error

		logFile, errorCreateFile := os.Create(path)

		if errorCreateFile != nil {
			log.Println("Canot create log file", errorCreateFile, logFile)
		}
	}

	return logFile, nil
}

func DefaultLogging(path ...string) {
	gin.DisableConsoleColor()

	if len(path) > 0 && path[0] != "" {
		defaultLogPath = path[0]
	} else {
		date := time.Now().Format("2006-01")
		defaultLogPath = "./logs/gin" + date + ".log"
	}

	createLogFolderIfNotExist(defaultLogPath)
	f, _ := openOrCreateLogFile(defaultLogPath)

	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mutex.Lock()

		// Determine the current filename
		date := time.Now().Format("2006-01-02")
		fileName := "logs/access-" + date + ".log"

		// If the filename has changed, update the log file
		if fileName != currentLogFile {
			if logger != nil && currentLogFile != "" {
				// Close the previous log file
				_ = os.Stdout.Close()
			}

			logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("Failed to create log file: %s", err)
			}

			// Create a new logger
			logger = log.New(logFile, "", log.LstdFlags)
			currentLogFile = fileName
		}

		mutex.Unlock()

		startTime := time.Now()

		// Process the request
		c.Next()

		// Log details after request is processed
		duration := time.Since(startTime).Milliseconds()
		label := "OK"
		if duration > 1000 {
			label = "SLOW_RESPONSE"
		} else if duration < 200 {
			label = "GREAT"
		}

		logger.Printf(
			"%s - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\" %d \"%s\"",
			c.ClientIP(),
			startTime.Format(time.RFC1123),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Writer.Status(),
			c.Writer.Size(),
			c.Request.Referer(),
			c.Request.UserAgent(),
			duration,
			label,
		)
	}
}
