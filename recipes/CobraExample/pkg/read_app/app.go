package read_app

import (
	"fmt"
	"os"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("cannot open %s\n", filename)
	}
	return string(data)
}

// Run the actual app
func Run(filename string) {
	fmt.Printf("Opening filename %s\n", filename)
	text := readFile(filename)
	fmt.Println(text)

}
