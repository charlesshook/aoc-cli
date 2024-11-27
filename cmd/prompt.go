package cmd

import (
	"fmt"
	"os"
	"time"

	utils "github.com/charlesshook/aoc-cli/internal/utils"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Gets AOC prompt.",
	Long:  "Gets the Advent of Code prompt for a given day and year",
	Run:   getPrompt,
}

func init() {
	rootCmd.AddCommand(promptCmd)

	promptCmd.Flags().IntP("day", "d", time.Now().Day(), "Day to grab the input file.")
	promptCmd.Flags().IntP("year", "y", time.Now().Year(), "Year to grab the input file.")
	promptCmd.Flags().StringP("cookie", "c", "", "Cookie used to make the request to Advent of Code.")
}

func getPrompt(cmd *cobra.Command, args []string) {
	day, _ := cmd.Flags().GetInt("day")
	year, _ := cmd.Flags().GetInt("year")
	cookie, _ := cmd.Flags().GetString("cookie")

	if day < 1 || day > 25 {
		fmt.Println("Day must be between 1 and 25!")
		return
	}

	if year < 2015 || year > time.Now().Year() {
		fmt.Printf("Year cannot be before 2015 or after %d.\n", time.Now().Year())
		return
	}

	if cookie == "" {
		fmt.Println("Cookie cannot be left empty")
		return
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day)
	data, _ := utils.GetDataFromUrlWithCookie(url, cookie)

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Could not get cwd")
		return
	}

	utils.WriteToFile(fmt.Sprintf("%s/%d/%d/PROMPT.md", cwd, year, day), data)
}
