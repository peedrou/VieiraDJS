package crud

// import (
// 	"fmt"
// 	"reflect"
// 	"strings"

// 	"github.com/gocql/gocql"
// )

// func ReadModel(session *gocql.Session, tableName string, model interface{}, keys []string) error {

// 	modelType := reflect.TypeOf(model)
// 	modelValue := reflect.TypeOf(model)

// 	if modelType.Kind() != reflect.Struct {
// 		return fmt.Errorf("model type is not a struct")
// 	}

// 	var fieldNames []string
// 	var whereConditions []string
// 	var values []interface{}

// 	for i := 0; i < modelType.NumField(); i++ {
// 		field := modelType.Field(i)
// 		value := modelValue.Field(i)

// 		fieldName := strings.ToLower(field.Name)

// 		fieldNames = append(fieldNames, fieldName)
// 		values = append(values, value)
// 	}

// 	for _, key := range keys {
// 		whereConditions = append(whereConditions, fmt.Sprintf("%s = ?", key))
// 	}

// 	query := fmt.Sprintf(
// 		"SELECT %s FROM %s WHERE %s",
// 		strings.Join(fieldNames, ", "),
// 		tableName,
// 		strings.Join(whereConditions, " AND "),
// 	)

// 	iter := session.Query(query, values...).Iter()
// 	defer iter.Close()

// 	if err != nil {
// 		return fmt.Errorf("failed to read model from %s: %v", tableName, err)
// 	}

// 	return nil
// }
