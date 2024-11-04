package crud

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"
)

func CreateModel(session *gocql.Session, tableName string, fieldNames []string, values ...interface{}) error {
	var placeholders []string
	len := len(values)

	for i := 0; i < len; i++ {
		placeholders = append(placeholders, "?")
	}
	// _, _ := DynamicModelBuilder(values)

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
