package getfromjson

import (
	"encoding/json"
	"os"
)

type DatabaseConnection struct {
	DATABASE_URL string
}

func GetDatabaseConData() (string, error) {
	data, err := os.ReadFile("../../configs/database.json")
	if err != nil {
		return "", err
	}

	var payload DatabaseConnection
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return "", err
	}

	return payload.DATABASE_URL, nil
}
