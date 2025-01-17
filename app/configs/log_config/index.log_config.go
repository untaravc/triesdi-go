package log_config

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	currentLogFile string
	logger         *log.Logger
	mutex          sync.Mutex
	defaultLogPath = "./logs/gin.log"
)

// Ensure the log folder exists
func createLogFolderIfNotExist(path string) {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755) // Use 0755 for directories
		if err != nil {
			log.Println("Failed to create directory:", err)
		} else {
			log.Println("Directory created:", dir)
		}
	}
}

// Open or create a log file
func openOrCreateLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

// Configure the default logging
func DefaultLogging(path ...string) {
	gin.DisableConsoleColor()

	if len(path) > 0 && path[0] != "" {
		defaultLogPath = path[0]
	} else {
		date := time.Now().Format("2006-01")
		defaultLogPath = "./logs/gin-" + date + ".log"
	}

	createLogFolderIfNotExist(defaultLogPath)
	f, _ := openOrCreateLogFile(defaultLogPath)

	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)
}

// Custom response writer to capture the response body
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Capture the response body
	return w.ResponseWriter.Write(b)
}

// Middleware for logging
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
				logger = nil
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

		// Capture the start time
		startTime := time.Now()

		// Read and capture the request body
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)
			// Reassign the body so it can be read by the handler
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Capture the response body
		writer := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// Process the request
		c.Next()

		// Calculate the duration of the request
		duration := time.Since(startTime).Milliseconds()

		// Determine the response label
		label := "OK"
		if duration > 1000 {
			label = "SLOW_RESPONSE"
		} else if duration < 200 {
			label = "GREAT"
		}

		// Log the details of the request and response
		logger.Printf(
			"\nRequest Time: %s\n"+
				"Client IP: %s\n"+
				"Method: %s\n"+
				"URL: %s\n"+
				"Response Status: %d\n"+
				"Request Duration: %d ms\n"+
				"Label: %s\n"+
				"Request Body: %s\n"+
				"Response Body: %s\n"+
				"User Agent: %s\n"+
				"Referer: %s\n",
			startTime.Format(time.RFC1123),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
			label,
			requestBody,
			writer.body.String(),
			c.Request.UserAgent(),
			c.Request.Referer(),
		)
	}
}
