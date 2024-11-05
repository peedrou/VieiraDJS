package crud

import (
	"fmt"

	"github.com/gocql/gocql"
)

func UpdateModelSingleKey(session *gocql.Session, tableName string, keyToUpdate string, valueToUpdate string, keyToSearch string, valueToSearch string) error {

	query := fmt.Sprintf(
		"UPDATE %s SET %s = '%s' WHERE %s = '%s'",
		tableName,
		keyToUpdate,
		valueToUpdate,
		keyToSearch,
		valueToSearch,
	)

	err := session.Query(query).Exec()

	if err != nil {
		return fmt.Errorf("failed to update model from %s: %v", tableName, err)
	}

	return nil
}

func UpdateModelBatch(session *gocql.Session, tableName string, keyToUpdate string, valueToUpdate string, keyToSearch string, IDs []interface{}) error {
	if len(IDs) == 0 {
		return fmt.Errorf("no IDs provided for updating")
	}

	for _, id := range IDs {
		query := fmt.Sprintf(
			"UPDATE %s SET %s = ? WHERE %s = ?",
			tableName,
			keyToUpdate,
			keyToSearch,
		)
		err := session.Query(query, valueToUpdate, id).Exec()
		if err != nil {
			return fmt.Errorf("failed to update model from %s for ID %s: %v", tableName, id, err)
		}
	}

	return nil
}
