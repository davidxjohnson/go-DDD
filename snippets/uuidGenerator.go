package main

import (
	"fmt"
	
	"github.com/satori/go.uuid"
)

func generateUUID() {
	guid := (uuid.Must(uuid.NewV4())).String()
	guidOutput := fmt.Sprintf("Your UUID is %s", guid)
	
	fmt.Println(guidOutput)
}

func main() {
	generateUUID()
}
