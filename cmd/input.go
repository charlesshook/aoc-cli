package cmd

import (
	"fmt"
	"os"
	"time"

	utils "github.com/charlesshook/aoc-cli/internal/utils"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Gets input file.",
	Long:  "Gets the Advent of Code input file for a given day and year",
	Run:   getInput,
}

func init() {
	rootCmd.AddCommand(inputCmd)

	inputCmd.Flags().IntP("day", "d", time.Now().Day(), "Day to grab the input file.")
	inputCmd.Flags().IntP("year", "y", time.Now().Year(), "Year to grab the input file.")
	inputCmd.Flags().StringP("cookie", "c", "", "Cookie used to make the request to Advent of Code.")
}

func getInput(cmd *cobra.Command, args []string) {
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

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	data, _ := utils.GetDataFromUrlWithCookie(url, cookie)

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Could not get cwd")
		return
	}

	utils.WriteToFile(fmt.Sprintf("%s/%d/%d/input.txt", cwd, year, day), data)
}
