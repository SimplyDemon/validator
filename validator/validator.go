package validator

import (
	"strings"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"log"
)

type Validator struct {
	Json             string
	PathToJsonSchema string
	IsJson           bool
}

func (validator *Validator) SetJSON(jsonText string) {
	validator.Json = jsonText
}

func (validator *Validator) CheckJson(text string) {
	trim := strings.TrimSpace(text)
	check := string(trim[0]) == "{" && string(trim[len(trim) - 1]) == "}"
	validator.IsJson = check
}

func (validator *Validator) SetJSONSchema(jsonSchemaPath string) {
	validator.PathToJsonSchema = jsonSchemaPath
}

func (validator *Validator) validate(jsonSchema string, jsonString string) bool {
	validator.CheckJson(jsonString)
	if !validator.IsJson {
		fmt.Println("Attention! Wrong JSON text")
	}
	schemaLoader := gojsonschema.NewReferenceLoader(jsonSchema)
	documentLoader := gojsonschema.NewStringLoader(jsonString)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false
		log.Fatal(err)
		panic(err.Error())
	}
	return result.Valid()
}