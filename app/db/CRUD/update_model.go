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