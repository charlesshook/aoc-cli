package languages

import "fmt"

func Setup(language string, day int, year int) {
	switch language {
	case "go":
		_ = setupGoProject(day, year)
	default:
		fmt.Printf("%s is not supported yet. Please feel free to open an issue on Github to request it.\n", language)
	}
}
