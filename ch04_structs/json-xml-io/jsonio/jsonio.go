package jsonio

import (
	"encoding/json"
	"os"
)

func LoadFromJson(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}

	return nil
}

func SaveToJson(output *os.File, key interface{}) error {
	encodeJSON := json.NewEncoder(output)
	err := encodeJSON.Encode(key)
	if err != nil {
		return err
	}

	return nil
}
