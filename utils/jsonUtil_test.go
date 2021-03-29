package utils

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
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
	result:=make(map[string]interface{})
	err := Unmarshal(jsonData, &result)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	fmt.Println(result["company"])
}

func TestJsonGet(t *testing.T) {
	result, _ :=jsonparser.GetString(jsonData,"company","name")
	assert.Equal(t, result,"Acme")

}
