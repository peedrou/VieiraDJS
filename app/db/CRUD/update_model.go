// package crud

// import (
// 	"fmt"
// 	"reflect"
// 	"strings"

// 	"github.com/gocql/gocql"
// )

// func UpdateModelSingleKey(session *gocql.Session, tableName string, key string, value string) error {

// 	query := fmt.Sprintf(
// 		"UPDATE %s",
// 		tableName,
// 		strings.Join(fieldNames, " AND "),
// 	)


// 	err := session.Query(query, values...).Exec()

// 	if err != nil {
// 		return fmt.Errorf("failed to remove model from %s: %v", tableName, err)
// 	}

// 	return nil
// }