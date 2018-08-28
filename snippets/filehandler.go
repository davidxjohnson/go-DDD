package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filepath string) []byte {
	fileContents, err := ioutil.ReadFile(filepath)
	
	if err != nil {
		fmt.Println(err)
	}
	
	return fileContents
}

func main() {
	filepath := "/tmp/helloWorld"
	fmt.Println(string(readFile(filepath)))
}
