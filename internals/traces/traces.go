package traces

import (
	"encoding/json"
	"os"
)

type StageGrade struct {
	StageMap string `json:"map_path"`
	Grade    int    `json:"grade"`
	Status   string `json:"status"`
}

type Traces struct {
	Compilation string       `json:"compilation"`
	Grades      []StageGrade `json:"grades"`
	FinalGrade  int          `json:"final_grade"`
}

type TracesInterface interface {
	AddCompilation(msg string)
	AddStage(mapName string, output int, status string)
	StoreInFile(path string) error
}

func NewLogger() *Traces {
	return &Traces{
		Compilation: "OK",
		Grades:      []StageGrade{},
		FinalGrade:  0,
	}
}

func (t *Traces) AddCompilation(msg string) {
	t.Compilation = msg
}

func (t *Traces) AddStage(mapName string, grade int, status string) {
	t.Grades = append(t.Grades, StageGrade{
		StageMap: mapName,
		Grade:    grade,
		Status:   status,
	})
	t.FinalGrade += grade
}

func (t *Traces) StoreInFile(file string) error {
	results, _ := json.Marshal(t)
	return os.WriteFile(file, results, 0644)
}
