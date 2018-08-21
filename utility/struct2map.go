import (
	"github.com/fatih/structs"
)

func convertStructToMap(conversionCandidate struct) map[string]interface{} {
	structs.Map(conversionCandidate)
