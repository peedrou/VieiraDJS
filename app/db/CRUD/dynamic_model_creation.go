package crud

import (
	"fmt"
	"reflect"
)

func DynamicModelBuilder(model ...interface{}) ([]string, []interface{}) {
	var placeholders []string
	var values []interface{}

	for range model {
		placeholders = append(placeholders, "?")
	}
	for _, m := range model {
		// Get the value of the current model
		modelValue := reflect.ValueOf(m)
		num := modelValue.NumField()
		fmt.Errorf("%s", num)

	}

	// modelValue := reflect.ValueOf(model)
	// modelType := modelValue.Type()

	// for i := 0; i < modelType.NumField(); i++ {
	// 	value := modelValue.Field(i)

	// 	placeholders = append(placeholders, "?")
	// 	values = append(values, value)
	// }
	return placeholders, values
}
