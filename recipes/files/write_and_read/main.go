package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
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

func zipFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("handle this error")
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err = gz.Write(data)

	if err != nil {
		fmt.Println("handle errer: ", err)
		return
	}

	if err = gz.Flush(); err != nil {
		return
	}

	if err = gz.Close(); err != nil {
		return
	}

	compressedData := b.Bytes()
	zippedName := fmt.Sprintf("%s_zipped", filename)
	fmt.Println("writing zipped data to :", zippedName)
	f, err := os.Create(zippedName)
	if err != nil {
		fmt.Println("\nshould handle this error: ", err)
		return
	}
	defer f.Close()
	f.Write(compressedData)

}

func unzipAndPrint(filename string) {
	fmt.Println("unzip file ", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("handle this error")
	}
	b := bytes.NewBuffer(data)

	var r io.Reader
	r, err = gzip.NewReader(b)
	if err != nil {
		return
	}

	var resB bytes.Buffer
	_, err = resB.ReadFrom(r)
	if err != nil {
		return
	}

	resData := resB.Bytes()
	fmt.Print(string(resData[0:100]))

}
func main() {
	filename := "/tmp/example.txt"
	createAndWriteToFile(filename)
	readTheEntireFile(filename)
	zipFile(filename)
	unzipAndPrint(fmt.Sprintf("%s_zipped", filename))
}
