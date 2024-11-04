package main

import (
	"fmt"
	"streamer"
)

func main() {
	// define number of workers and jobs
	const numJobs = 4
	const numWorkers = 2

	// create channels for work and result:
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool.
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println(wp)
	// Start the worker pool

	// Create 4 videos to send to the worker pool

	// Send the videos to the worker pool

	// Print the results
}
