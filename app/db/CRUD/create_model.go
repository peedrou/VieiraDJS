package crud

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gocql/gocql"
)

func CreateModel(session *gocql.Session, tableName string, model interface{}) error {

	modelType := reflect.TypeOf(model)
	modelValue := reflect.TypeOf(model)

	if modelType.Kind() != reflect.Struct {
		return fmt.Errorf("model type is not a struct")
	}

	fieldNames, placeholders, values := DynamicModelBuilder(modelType, modelValue)

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(fieldNames, ", "),
		strings.Join(placeholders, ", "),
	)

	err := session.Query(query, values...).Exec()
	if err != nil {
		return fmt.Errorf("failed to insert model into %s: %v", tableName, err)
	}

	return nil

}

