package helpers

import "encoding/json"

// decoding json
func DecodeJson(data []byte, decodedJson any) {
	var err = json.Unmarshal(data, decodedJson)
	if err != nil {
		panic(err)
	}
}

// encoding to json
func EncodeJson(data any) []byte {
	json, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	return json
}
