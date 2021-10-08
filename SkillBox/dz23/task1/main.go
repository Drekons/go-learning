package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		file1      string
		file2      string
		resultFile string
	)

	if len(os.Args) > 1 {
		file1 = os.Args[1]
	}
	if len(os.Args) > 2 {
		file2 = os.Args[2]
	}
	if len(os.Args) > 3 {
		resultFile = os.Args[3]
	}

	joinedFilesBytes := joinFiles(file1, file2, resultFile)

	if joinedFilesBytes != nil {
		fmt.Printf("%s\n", string(joinedFilesBytes))
	}
}

func joinFiles(file1, file2, resultFile string) []byte {
	if file1 == "" {
		return nil
	}
	file1Bytes := readFile(file1)

	if file2 == "" {
		return file1Bytes
	}
	file2Bytes := readFile(file2)

	joinedFiles := strings.Join([]string{string(file1Bytes), string(file2Bytes)}, "\n")
	joinedFilesBytes := []byte(joinedFiles)

	if resultFile == "" {
		return joinedFilesBytes
	}

	err := ioutil.WriteFile(getFilesDir()+resultFile, joinedFilesBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func readFile(fileName string) []byte {
	readBytes, err := ioutil.ReadFile(getFilesDir() + fileName)
	if err != nil {
		panic(err)
	}

	return readBytes
}

func getFilesDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return pwd + "/files/"
}
