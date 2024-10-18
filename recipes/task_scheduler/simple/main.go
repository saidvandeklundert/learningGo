package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func periodicCommandStdoutAndStderr() {
	fmt.Println("periodicCommandStdoutAndStderr job 4:")
	// Define the command to run (e.g., 'ls -l /nonexistent')
	cmd := exec.Command("ls", "-l", "/nonexistent")

	// Buffers to capture stdout and stderr
	var stdoutBuf, stderrBuf bytes.Buffer

	// Redirect the stdout and stderr to buffers
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	// Run the command
	if err := cmd.Run(); err != nil {
		log.Printf("\tCommand execution error: %v", err)
	}

	// Print the captured stdout and stderr
	fmt.Printf("\tStdout:\n%s\n", stdoutBuf.String())
	fmt.Printf("\tStderr:\n%s\n", stderrBuf.String())
}

func periodicCommand() {
	fmt.Println("periodicCommand job 3:")
	// Define the command you want to execute (e.g., 'ls -l')
	cmd := exec.Command("ls", "-l")

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	// Print the output as a string
	fmt.Println("\t", string(output))
}

func main() {
	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
	}

	// add a job to the scheduler
	job1, err := s.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println("anonfunc job 1")
			},
			"hello",
			1,
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	// add another job:
	job2, err := s.NewJob(
		gocron.DurationJob(
			1*time.Second,
		),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println("anonfunc job 2")
				fmt.Println("just doing the thing that needs doing", a, b)

			},
			"hello",
			1,
		),
	)

	if err != nil {
		fmt.Println(err)
	}
	// add another job:
	job3, err := s.NewJob(
		gocron.DurationJob(
			5*time.Second,
		),
		gocron.NewTask(
			periodicCommand,
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	// add another job:
	job4, err := s.NewJob(
		gocron.DurationJob(
			5*time.Second,
		),
		gocron.NewTask(
			periodicCommandStdoutAndStderr,
		),
	)

	if err != nil {
		fmt.Println(err)
	}
	// each job has a unique id
	fmt.Println(job1.ID())
	fmt.Println(job2.ID())
	fmt.Println(job3.ID())
	fmt.Println(job4.ID())
	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute):
	}

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		// handle error
	}

}
