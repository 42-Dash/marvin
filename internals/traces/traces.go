package traces

import (
	"fmt"
	"os"
)

type Traces struct {
	File *os.File
}

type TracesInterface interface {
	CloseLogger() error

	CompilationError(msg string) error
	CompilationSuccess() error

	GradingError(mapName string, output int, reason string) error
	TimeoutError(mapName, output string) error
	UploadMap(mapName string) error

	GradingSuccess(input, output string) error
}

func NewLogger(path string) (*Traces, error) {
	file, err := os.Create(path)
	if err != nil {
		return &Traces{}, err
	}
	return &Traces{File: file}, nil
}

func (t *Traces) CloseLogger() error {
	return t.File.Close()
}

func (t *Traces) CompilationError(msg string) error {
	_, err := t.File.WriteString("Compilation error: " + msg + "\n")
	return err
}

func (t *Traces) CompilationSuccess() error {
	_, err := t.File.WriteString("Compilation success\n")
	return err
}

func (t *Traces) TimeoutError(mapName string) error {
	_, err := t.File.WriteString("Time limit exceeded for map: " + mapName + "\n")
	return err
}

func (t *Traces) UploadMap(mapName string) error {
	_, err := t.File.WriteString("Uploading map: " + mapName + "\n")
	return err
}

func (t *Traces) GradingError(mapName string, output int, reason string) error {
	_, err := t.File.WriteString(fmt.Sprintf("Error grading map: %v\nMarvin's output: %v\nReason: %v\n", mapName, output, reason))
	return err
}

func (t *Traces) GradingSuccess(input string, output int) error {
	_, err := t.File.WriteString(fmt.Sprintf("Success: %v, final grade: %v", input, output))
	return err
}
