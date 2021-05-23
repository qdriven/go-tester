package json

import (
	"bytes"
	stdjson "encoding/json"
	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

var (
	JSON = jsoniter.Config{
		EscapeHTML:             false,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
)

var (
	arrayPrefix  = []byte("[")
	objectPrefix = []byte("{")
)

type RawMessage = stdjson.RawMessage

// UseNumber solve very big int64 digits loss.
func UseNumber() {
	JSON = jsoniter.Config{
		UseNumber:              true,
		EscapeHTML:             false,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
}

// Marshal returns the JSON encoding of v.
func Marshal(v interface{}) ([]byte, error) {
	return JSON.Marshal(v)
}

// MustMarshal must returns the JSON encoding of v.
func MustMarshal(v interface{}) []byte {
	data, _ := JSON.Marshal(v)
	return data
}

// MarshalToString returns the JSON encoding to string of v.
func MarshalToString(v interface{}) (string, error) {
	return JSON.MarshalToString(v)
}

// MustMarshalToString must returns the JSON encoding to string of v.
func MustMarshalToString(v interface{}) string {
	str, _ := JSON.MarshalToString(v)
	return str
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	return JSON.Unmarshal(data, v)
}

func UnmarshalToMap(jsonStr string) (map[string]interface{}, error) {
	v := make(map[string]interface{})
	err :=UnmarshalFromString(jsonStr, &v)
	if err!=nil{
		return nil,err
	}
	return v,nil

}

// UnmarshalFromString unmarshal string to v.
func UnmarshalFromString(str string, v interface{}) error {
	return JSON.UnmarshalFromString(str, v)
}

// Valid check JSON data.
func Valid(data []byte) bool {
	return JSON.Valid(data)
}

// ValidFromString check JSON string.
func ValidFromString(str string) bool {
	return Valid([]byte(str))
}

/**
Get String from Json
for usage from:
https://github.com/buger/jsonparser
*/
func GetString(data []byte, path ...string) (string, error) {
	return jsonparser.GetString(data, path...)
}

/**
JsonPath in Golang usage
https://github.com/ohler55/ojg
*/
func GetByPath(str string, jsonPathExp string) ([]interface{}, error) {

	obj, err := oj.ParseString(str)
	if err != nil {
		return nil, err
	}
	x, err := jp.ParseString(jsonPathExp)
	if err != nil {
		return nil, err
	}
	ys := x.Get(obj)
	return ys, err
}

func Format(raw []byte) []byte {
	raw = bytes.TrimSpace(raw)
	if bytes.HasPrefix(raw, arrayPrefix) {
		var val []interface{}
		if err := Unmarshal(raw, &val); err != nil {
			return raw
		}

		data, err := Marshal(val)
		if err != nil {
			return raw
		}

		return data
	}

	if bytes.HasPrefix(raw, objectPrefix) {
		val := map[string]interface{}{}
		if err := Unmarshal(raw, &val); err != nil {
			return raw
		}

		data, err := Marshal(val)
		if err != nil {
			return raw
		}

		return data
	}

	return raw
}

func FormatFromString(raw string) string {
	return string(Format([]byte(raw)))
}
