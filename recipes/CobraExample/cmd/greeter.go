/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Name string

// greeterCmd represents the greeter command
var greeterCmd = &cobra.Command{
	Use:   "greeter",
	Short: "Say hi to target person",
	Long: `Say hi to target person.

Just says 'hi' to someone. Could not think of anything better.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("greeter called")
	},
}

func init() {
	rootCmd.AddCommand(greeterCmd)

	greeterCmd.Flags().StringVarP(&Name, "name", "n", "", "the name")
	greeterCmd.MarkFlagRequired("name")

}
