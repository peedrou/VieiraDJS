package crud

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"
)

func RemoveModel(session *gocql.Session, tableName string, IDField string, IDs []interface{}) error {
	if len(IDs) == 0 {
		return fmt.Errorf("no IDs provided for deletion")
	}

	placeholders := make([]string, len(IDs))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	// values := make([]interface{}, len(IDs))
	// for i, id := range IDs {
	// 	values[i] = id
	// }

	query := fmt.Sprintf(
		"DELETE FROM %s WHERE %s IN (%s)",
		tableName,
		IDField,
		strings.Join(placeholders, ", "),
	)

	err := session.Query(query, IDs...).Exec()

	if err != nil {
		return fmt.Errorf("failed to remove model from %s: %v", tableName, err)
	}

	return nil
}
