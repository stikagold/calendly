package helpers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func ConvertToBytes(o interface{}) []byte {
	resp, err := json.Marshal(o)
	if err != nil {
		HandleError(err)
	}
	return resp
}

func ConvertToHash(str string) (string, error) {
	resp, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	return string(resp), err
}

func GetFormattedUrl(prefix string, path string) string {
	if prefix == "" {
		return fmt.Sprintf("/%s", path)
	}

	return fmt.Sprintf("/%s/%s", prefix, path)
}
