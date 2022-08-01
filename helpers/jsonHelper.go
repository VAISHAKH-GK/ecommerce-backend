package helpers

import "encoding/json"

func DecodeJson(data []byte, decodedjson any) {
	var err = json.Unmarshal(data, &decodedjson)
	if err != nil {
		panic(err)
	}
}

func EncodeJson(data any) []byte {
	json, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	return json
}
