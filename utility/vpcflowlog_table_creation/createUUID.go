package main 

import (
	"fmt"

	"github.com/satori/go.uuid"
)

func generateUUID() uuid.UUID {
	uuid := uuid.Must(uuid.NewV4())
	return uuid
}

func main() {
	fmt.Println(generateUUID())
}
