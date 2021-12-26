package util

import (
	"log"

	json "github.com/json-iterator/go"
)

func JSONDump(v interface{}) string {
	s, err := json.MarshalToString(v)
	if err != nil {
		log.Println("json dump fail:", err)
		return ""
	}
	return s
}
