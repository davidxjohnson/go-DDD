package utility

import (
	"fmt"
	
	"github.com/satori/go.uuid"
)

func GenerateUUID() string {
	guid := (uuid.Must(uuid.NewV4())).String()
	guidOutput := fmt.Sprintf("Your UUID is %s", guid)
	
	return guidOutput
}
