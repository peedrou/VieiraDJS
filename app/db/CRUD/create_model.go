package crud

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"
)

func CreateModel(session *gocql.Session, tableName string, fieldNames []string, values []interface{}) error {

	// modelType := reflect.TypeOf(model)
	// modelValue := reflect.TypeOf(model)

	// if modelType.Kind() != reflect.Struct {
	// 	return fmt.Errorf("model type is not a struct")
	// }

	placeholders, values := DynamicModelBuilder(values)

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
