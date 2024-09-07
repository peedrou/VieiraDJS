package crud

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gocql/gocql"
)

func RemoveModel(session *gocql.Session, tableName string, model interface{}) error {

	modelType := reflect.TypeOf(model)
	modelValue := reflect.TypeOf(model)

	if modelType.Kind() != reflect.Struct {
		return fmt.Errorf("model type is not a struct")
	}

	fieldNames, _, values := DynamicModelBuilder(modelType, modelValue)

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s",
		tableName,
		strings.Join(fieldNames, " AND "),
	)


	err := session.Query(query, values...).Exec()

	if err != nil {
		return fmt.Errorf("failed to remove model from %s: %v", tableName, err)
	}

	return nil
}