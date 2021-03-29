package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFileInLines(filePath string) []string {
	byteContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(byteContents), "\n")
	return lines
}

func ReadJsonFile(filePath string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	byteContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	err = Unmarshal(byteContents, result)
	return result, err
}

func ReadYamlFile(filePath string) {

}
