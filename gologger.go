package gologger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// SetupLogger is a basic logger that creates and writes to a select file
func SetupLogger(filename string) *log.Logger {
	file := setupLogFile(filename)
	multi := io.MultiWriter(file, os.Stdout)
	Logger := log.New(multi, "LOGGER: ", log.Ldate|log.Ltime)
	return Logger
}

func setupLogFile(filename string) *os.File {
	_, err := os.Stat(filename) // does file exist?
	if err == nil {
		removeFile(filename)
	}

	file := createFile(filename)

	return file
}

func removeFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Fprintln(os.Stdout, "Failed to remove log file :", err)
		panic(err)
	}
}

func createFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Fprintln(os.Stdout, "Failed to create log file :", err)
		panic(err)
	}

	return file
}
