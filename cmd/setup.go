package cmd

import (
	"fmt"
	"time"

	languages "github.com/charlesshook/aoc-cli/internal/languages"
	"github.com/spf13/cobra"
)

var setupProjectCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up a project for a given day of AOC.",
	Long:  "Sets up a project for a given day of AOC.",
	Run:   setupProject,
}

func init() {
	rootCmd.AddCommand(setupProjectCmd)

	setupProjectCmd.Flags().IntP("day", "d", time.Now().Day(), "Day to create the project for.")
	setupProjectCmd.Flags().IntP("year", "y", time.Now().Year(), "Year to create the project for.")
	setupProjectCmd.Flags().StringP("language", "l", "go", "The language to create the project for")
}

func setupProject(cmd *cobra.Command, args []string) {
	day, _ := cmd.Flags().GetInt("day")
	year, _ := cmd.Flags().GetInt("year")
	language, _ := cmd.Flags().GetString("language")

	if day < 1 || day > 25 {
		fmt.Println("Day must be between 1 and 25!")
		return
	}

	if year < 2015 || year > time.Now().Year() {
		fmt.Printf("Year cannot be before 2015 or after %d.\n", time.Now().Year())
		return
	}

	languages.Setup(language, day, year)
}
