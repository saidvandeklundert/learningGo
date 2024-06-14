package main

import "github.com/saidvandeklundert/toolkit"

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}