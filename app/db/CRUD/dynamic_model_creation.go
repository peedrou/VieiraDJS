package crud

import (
	"reflect"
)

func DynamicModelBuilder(model []interface{}) ([]string, []interface{}) {
	var placeholders []string
	var values []interface{}

	modelValue := reflect.ValueOf(model)
	modelType := reflect.TypeOf(model)

	for i := 0; i < modelType.NumField(); i++ {
		value := modelValue.Field(i)

		placeholders = append(placeholders, "?")
		values = append(values, value)
	}
	return placeholders, values
}
