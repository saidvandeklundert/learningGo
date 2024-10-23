package main

import (
	"fmt"
	"os"
)

type Configuration struct {
	Parameter1 string
	Parameter2 string
	Parameter3 string
}

var config Configuration

// init functions are run without being called explicitly
// one usecase could be setting up a complex configuration of
// multiple objects
func init() {
	fmt.Println("init is running now!")
	environment := os.Getenv("AWS_LAMBDA")
	if environment != "" {
		config.Parameter1 = environment
	} else {
		config.Parameter1 = "NOT_ON_LAMBA"
	}

	config.Parameter2 = "something else"
	config.Parameter3 = "something special"
}

func main() {
	fmt.Println(config)
	fmt.Println("main coming to a close")

}
