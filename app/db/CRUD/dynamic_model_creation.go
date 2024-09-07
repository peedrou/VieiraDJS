package crud

import (
	"reflect"
	"strings"
)

func DynamicModelBuilder(modelType reflect.Type, modelValue reflect.Type) ([]string, []string, []interface{}) {
	var fieldNames []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		value := modelValue.Field(i)

		fieldName := strings.ToLower(field.Name)

		fieldNames = append(fieldNames, fieldName)
		placeholders = append(placeholders, "?")
		values = append(values, value)
	}
	return fieldNames, placeholders, values
}