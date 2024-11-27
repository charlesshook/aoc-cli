package languages

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	utils "github.com/charlesshook/aoc-cli/internal/utils"
)

//go:embed templates/go/main.go.txt
var mainGoFile []byte

//go:embed templates/go/test.go.txt
var testGoFile []byte

func setupGoProject(day int, year int) error {
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	pathToProject := fmt.Sprintf("%s/%d/%d/go/", cwd, year, day)

	_ = os.MkdirAll(filepath.Dir(pathToProject), os.ModePerm)

	utils.WriteToFile(fmt.Sprintf("%s/%s", filepath.Dir(pathToProject), "main.go"), mainGoFile)
	utils.WriteToFile(fmt.Sprintf("%s/%s", filepath.Dir(pathToProject), "test.go"), testGoFile)

	return nil
}
