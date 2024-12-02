package languages

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	utils "github.com/charlesshook/aoc-cli/internal/utils"
)

//go:embed templates/python/main.py.txt
var mainPythonFile []byte

//go:embed templates/python/part_one.py.txt
var partOnePythonFile []byte

//go:embed templates/python/part_two.py.txt
var partTwoPythonFile []byte

func setupPythonProject(day int, year int) error {
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	pathToProject := fmt.Sprintf("%s/%d/%d/python/", cwd, year, day)

	_ = os.MkdirAll(filepath.Dir(pathToProject), os.ModePerm)

	utils.WriteToFile(fmt.Sprintf("%s/%s", filepath.Dir(pathToProject), "main.py"), mainPythonFile)
	utils.WriteToFile(fmt.Sprintf("%s/%s", filepath.Dir(pathToProject), "part_one.py"), partOnePythonFile)
	utils.WriteToFile(fmt.Sprintf("%s/%s", filepath.Dir(pathToProject), "part_two.py"), partTwoPythonFile)

	return nil
}
