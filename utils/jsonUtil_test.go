package utils

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var (
	jsonData = []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)
	jsonString = `{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`
)

func TestJsonUnMarsh(t *testing.T) {
	result := make(map[string]interface{})
	err := Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(result["company"])
}

func TestJsonGet(t *testing.T) {
	result, _ := jsonparser.GetString(jsonData, "company", "name")
	assert.Equal(t, result, "Acme")

}

func test(value interface{}) {
	switch value.(type) {
	case string:
		// 将interface转为string字符串类型
		op, ok := value.(string)
		fmt.Println(op, ok)
	case int32:
		// 将interface转为int32类型
		op, ok := value.(int32)
		fmt.Println(op, ok)
	case int64:
		// 将interface转为int64类型
		op, ok := value.(int64)
		fmt.Println(op, ok)
	case []int:
		// 将interface转为切片类型
		op := make([]int, 0)
		op = value.([]int)
		fmt.Println(op)
	default:
		fmt.Println("unknown")
	}
}

func TestGoGetByPath(t *testing.T) {
	result, err := GetByPath(jsonString, "company.name")
	if err != nil {
		assert.True(t, false)
	}
	fmt.Println(result[0])
	assert.Equal(t, result[0], "Acme")
	fmt.Println(reflect.TypeOf(result[0]))
}
