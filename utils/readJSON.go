package utils

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadJSON(filename string, v interface{}) error {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return errors.New("o arquivo não existe: " + filename)
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(v)
}
