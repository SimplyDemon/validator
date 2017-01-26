package validator

import (
	"strings"
	"github.com/xeipuuv/gojsonschema"
)

type Validator struct {
	Json             string
	PathToJsonSchema string
}

func (validator *Validator) SetJSON(jsonText string) {
	validator.Json = jsonText
}

func (validator *Validator) CheckJson(text string) bool {
	trim := strings.TrimSpace(text)
	check := string(trim[0]) == "{" && string(trim[len(trim) - 1]) == "}"
	return check
}

func (validator *Validator) SetJSONSchema(jsonSchemaPath string) {
	validator.PathToJsonSchema = jsonSchemaPath
}

func (validator *Validator) IsValid() bool {
	if !validator.CheckJson(validator.Json) {
		return false
	}
	schemaLoader := gojsonschema.NewReferenceLoader(validator.PathToJsonSchema)
	documentLoader := gojsonschema.NewStringLoader(validator.Json)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return false
	}
	return result.Valid()
}