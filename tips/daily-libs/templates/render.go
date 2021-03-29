package templates

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"unicode"
)

/**
1. Go-Template Render
2. JSON/YAML data structure convertion
3. Go File Reading/Writing
4. Go CLI Usage
5. RENDER Pipeline
*/

type FileTemplateContext struct {
	TemplateFile    string
	ContextDataFile string
	ContextData     interface{}
	OutDir          string
	OutPutFile      string
}

var (
	jsonPrefix = []byte("{")
)

func isJSON(buf []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, jsonPrefix)
}

func ToJSON(data []byte) ([]byte, error) {
	if isJSON(data) {
		return data, nil
	}
	return yaml.YAMLToJSON(data)
}

func Render(ctx *FileTemplateContext) {
	//setup template,todo: regist new function
	template, err := template.New(path.Base(ctx.TemplateFile)).Funcs(RenderFuncMapping).ParseFiles(ctx.TemplateFile)
	if err != nil {
		fmt.Println(err)
	}
	var dataInFile interface{}
	if ctx.ContextDataFile != "" {
		dataFile, ioErr := os.Open(ctx.ContextDataFile)
		if ioErr != nil {
			fmt.Println(ioErr)
		}
		byteValue, err := ioutil.ReadAll(dataFile)
		if err != nil {
			fmt.Println(err)
		}
		byteValue, err = ToJSON(byteValue)
		err = json.Unmarshal(byteValue, &dataInFile)
		if err != nil {
			fmt.Println(err)
		}
		defer dataFile.Close()
	}
	//FILE or Data
	if dataInFile != nil {
		RenderFile(template, ctx.OutDir, ctx.OutPutFile, dataInFile)
	} else {
		RenderFile(template, ctx.OutDir, ctx.OutPutFile, ctx.ContextData)
	}
}

func RenderFile(template *template.Template, outDir, outputFileName string, contextData interface{}) {
	absOutputFileName := path.Join(outDir, outputFileName)
	os.MkdirAll(path.Dir(absOutputFileName), os.ModePerm)
	outputFile, err := os.Create(absOutputFileName)
	defer outputFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = template.Execute(outputFile, contextData)
	if err != nil {
		fmt.Println(err)
	}
}
