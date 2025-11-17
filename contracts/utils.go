package contracts

import (
	"encoding/json"

	"github.com/alpkeskin/gotoon"
)

func ToJson(data interface{}) string {
	json, err := json.Marshal(data)

	if err != nil {
		return ""
	}

	return string(json)
}

func ToToon(data interface{}) string {
	json, err := gotoon.Encode(data)

	if err != nil {
		return ""
	}

	return string(json)
}
