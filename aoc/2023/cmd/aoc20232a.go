/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

//go run . aoc20232a
import (
	"aoc/pkg/test"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type GameRound struct {
	Red   int
	Blue  int
	Green int
}

type Game struct {
	Number int
	line   string
	Red    int
	Blue   int
	Green  int
	Rounds []GameRound
}

func NewGame(line string) Game {
	var game Game
	game.Rounds = []GameRound{}

	game.line = line
	return game
}

func (g *Game) SetNumber() {
	chunks := strings.Split(g.line, ":")

	match_number := strings.Split(chunks[0], " ")[1]
	g.Number, _ = strconv.Atoi(match_number)

}

// Check if a game is possible given some max values
func (g *Game) IsPossible(game Game, maxRed, maxGreen, maxBlue int) bool {
	for _, round := range game.Rounds {
		if round.Red > maxRed || round.Green > maxGreen || round.Blue > maxBlue {

			return false

		}
	}
	return true
}
func (g *Game) SetRoundsAndGameTotal() {
	chunks := strings.Split(g.line, ":")
	game_number := strings.Split(chunks[1], ";")
	for _, round := range game_number {
		var gameRound GameRound
		for _, turn := range strings.Split(round, ",") {

			turn_result := strings.Split(turn, " ")
			turn_color := turn_result[2]
			turn_amount := turn_result[1]
			turn_amount_int, _ := strconv.Atoi(turn_amount)

			if turn_color == "red" {
				g.Red += turn_amount_int
				gameRound.Red = turn_amount_int
			} else if turn_color == "green" {
				g.Green += turn_amount_int
				gameRound.Green = turn_amount_int
			} else if turn_color == "blue" {
				g.Blue += turn_amount_int
				gameRound.Blue = turn_amount_int
			}

		}
		g.Rounds = append(g.Rounds, gameRound)
	}
}

func GetGame(line string) Game {

	game := NewGame(line)
	game.SetNumber()
	game.SetRoundsAndGameTotal()
	return game
}

func GetGames(assignmentInput string) []Game {
	var games []Game
	for _, line := range strings.Split(assignmentInput, "\n") {
		game := GetGame(line)
		games = append(games, game)
	}
	return games

}

func Puzzle20232a() {

	assignmentInput := test.GetPuzzleInput("2023", "2")

	fmt.Println("Puzzle20232a ran")

	games := GetGames(assignmentInput)

	solution := 0
	for _, game := range games {
		//only 12 red cubes, 13 green cubes, and 14 blue cubes
		if game.IsPossible(game, 12, 13, 14) {
			solution += game.Number
		}
	}

	fmt.Println(solution)

	expected := 2795
	test.Assert(solution == expected, fmt.Sprintf("got %d want %d", solution, expected))

}

// aoc20232aCmd represents the aoc20232a command
var aoc20232aCmd = &cobra.Command{
	Use:   "aoc20232a",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aoc20232a called")
		Puzzle20232a()
	},
}

func init() {
	rootCmd.AddCommand(aoc20232aCmd)

}
