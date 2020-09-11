package xmlio

import (
	"encoding/xml"
	"os"
)

func LoadFromXml(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	decodeJSON := xml.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}

	return nil
}

func SaveToXml(output *os.File, key interface{}) error {
	encodeJSON := xml.NewEncoder(output)
	err := encodeJSON.Encode(key)
	if err != nil {
		return err
	}

	return nil
}
