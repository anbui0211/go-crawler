package services

import (
	"encoding/json"
	"log"
	"os"
)



func WriteToJSONFile(fileName string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("Could not marshal data to JSON: %v", err)
	}

	return os.WriteFile(fileName, jsonData, 0644)
}
