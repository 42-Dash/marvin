package logger

import (
	"log"
	"os"
	"sync"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	file  *os.File
	once  sync.Once
	mu    sync.Mutex
)

func InitLogger() error {
	var err error
	file, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// Create loggers that write to the file
	Info = log.New(file, "[INFO]: ", log.Ldate|log.Ltime)
	Error = log.New(file, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(file, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func Flush() error {
	mu.Lock()
	defer mu.Unlock()
	if file != nil {
		return file.Sync()
	}
	return nil
}

func CloseFile() {
	once.Do(func() {
		if file != nil {
			err := file.Close()
			if err != nil {
				log.Fatalf("Failed to close log file: %v", err)
			}
		}
	})
}
