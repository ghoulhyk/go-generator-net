package jsonutil

import gojson "github.com/goccy/go-json"

func ToJson(v any) (string, error) {
	bytes, err := gojson.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ToByte(v any) ([]byte, error) {
	bytes, err := gojson.Marshal(v)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

func ToJsonOrEmpty(v any) (json string) {
	json, _ = ToJson(v)
	return json
}

func ToByteOrEmpty(v any) (json []byte) {
	json, _ = ToByte(v)
	return json
}

func FromJson(data string, v any) error {
	return FromByte([]byte(data), v)
}

func FromByte(data []byte, v any) error {
	return gojson.Unmarshal(data, v)
}
