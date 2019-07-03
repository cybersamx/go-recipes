package utils

import (
	"encoding/json"
	"net/http"
)

type (
	Event struct {
		Event string `json:"event"`
	}

	Messasge struct {
		Message string `json:"message"`
	}
)

func WriteJSONResponse(w http.ResponseWriter, payload interface{}) error {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	jsonData, err := json.Marshal(payload)
	if err == nil {
		w.Write(jsonData)
	}

	return err
}