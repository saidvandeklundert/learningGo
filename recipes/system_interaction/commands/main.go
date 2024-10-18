package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

type CommandResponse struct {
	Command []string
	Stdout  string
	Stderr  string
}

/*
Runs a given command the local system and returns STDOUT as
well as STDERR
*/
func sendSingleCommand(command ...string) (CommandResponse, error) {

	if len(command) == 0 {
		return CommandResponse{}, errors.New("this is a basic error")
	}
	var stdoutBuf, stderrBuf bytes.Buffer
	if len(command) == 1 {
		cmd := exec.Command(command[0])
		cmd.Stdout = &stdoutBuf
		cmd.Stderr = &stderrBuf
		if err := cmd.Run(); err != nil {
			return CommandResponse{}, err
		}
	} else {
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Stdout = &stdoutBuf
		cmd.Stderr = &stderrBuf
		if err := cmd.Run(); err != nil {
			return CommandResponse{}, err
		}
	}

	fmt.Printf("\tStdout:\n%s\n", stdoutBuf.String())
	fmt.Printf("\tStderr:\n%s\n", stderrBuf.String())
	response := CommandResponse{
		Stdout:  stdoutBuf.String(),
		Stderr:  stderrBuf.String(),
		Command: command,
	}
	return response, nil
}

func main() {
	fmt.Println("Examples on how to send commands from a program running on a Linux system.")

	commandResponse, err := sendSingleCommand("ls", "-ltr")
	if err != nil {
		fmt.Println("got an error ", err)
		return
	}
	fmt.Println(commandResponse.Stdout, commandResponse.Stderr, commandResponse.Command)

	commandResponse, err = sendSingleCommand("ls")
	if err != nil {
		fmt.Println("got an error ", err)
		return
	}
	fmt.Println(commandResponse.Stdout, commandResponse.Stderr, commandResponse.Command)

}
