package libraries

import (
	"encoding/json"
	"log"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SprintResponse(status string, message string) []byte {
	e := Response{
		Status:  status,
		Message: message,
	}

	bytes, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		log.Printf("%s", err)
	}

	return bytes
}
