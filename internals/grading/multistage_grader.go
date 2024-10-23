package grading

import (
	"dashinette/pkg/parser"
	"os"
)

func MultistageGrader(config parser.TesterConfig) (err error) {
	if _, err = os.Stat(config.Tracesfile); err == nil {
		os.Remove(config.Tracesfile)
	}

	traces, err := os.Create(config.Tracesfile)
	if err != nil {
		return
	}
	defer traces.Close()

	_, err = traces.Write([]byte(config.Repo + "\n" + config.Maps[0].Name + "\n" + config.Maps[0].Path + "\n"))

	return
}
