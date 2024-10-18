package main

import (
	"fmt"
	"os"
)

func createAndWriteToFile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("should handle this error: ", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("just putting something into a file\nnothing special")
	if err != nil {
		fmt.Println("should handle this error: ", err)
		return
	}
	for i := range 200000 {
		line := fmt.Sprintf("additional line %d that needs to be written to the file\n", i)
		_, err = f.WriteString(line)
		if err != nil {
			fmt.Println("should handle this error: ", err)
			return
		}
	}
	f.Sync()
}

func readTheEntireFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("handle this error")
	}

	fmt.Print(string(data[0:100]))

}
func main() {
	filename := "/tmp/example.txt"
	createAndWriteToFile(filename)
	readTheEntireFile(filename)
}
