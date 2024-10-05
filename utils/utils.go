package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		log.Fatal("Bad request")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}
