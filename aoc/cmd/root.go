/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Collection of Advent of Code solutions",
	Long: `This is a collection of solutions to several
Advent of Code challenges.

You can run the solution to any of the problems like this:

go run . <year><day><assignment:a|b>

So to run day 1, assignment b for year 2023
go run . 20231b `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
