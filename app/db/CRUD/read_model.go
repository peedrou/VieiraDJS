package crud

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"
)

func ReadModel(session *gocql.Session, tableName string, fields []string, keys []string, values ...interface{}) ([]interface{}, error) {

	if len(keys) != len(values) {
		return nil, fmt.Errorf("not all keys for the query have corresponding values")
	}

	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields specified to retrieve")
	}

	var whereConditions []string

	for _, key := range keys {
		whereConditions = append(whereConditions, fmt.Sprintf("%s = ?", key))
	}

	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s ALLOW FILTERING",
		strings.Join(fields, ", "),
		tableName,
		strings.Join(whereConditions, " AND "),
	)

	iter := session.Query(query, values...).Iter()
	defer iter.Close()

	var results []map[string]interface{}
	row := make(map[string]interface{})

	for iter.MapScan(row) {
		rowCopy := make(map[string]interface{}, len(row))

		for k, v := range row {
			rowCopy[k] = v
		}

		results = append(results, rowCopy)

		for k := range row {
			delete(row, k)
		}
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("failed to read model from %s: %v", tableName, err)
	}

	finalResults := getInterfacesFromMapSlice(results)

	return finalResults, nil
}

func getInterfacesFromMapSlice(maps []map[string]interface{}) []interface{} {
	var result []interface{}
	for _, m := range maps {
		for _, v := range m {
			result = append(result, v)
		}
	}
	return result
}
