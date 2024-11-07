package mocks

// import (
// 	"VieiraDJS/app/interfaces/gocql_interface"

// 	"github.com/Agent-Plus/gocqlmock"
// )

// type MockSession struct {
// 	*gocqlmock.Session
// }

// type MockQuery struct {
// 	*gocqlmock.Query
// }

// type MockIter struct {
// 	*gocqlmock.Iter
// }

// func (s MockSession) Query(query string, args ...interface{}) gocql_interface.QueryInterface {
// 	return &MockQuery{s.Session.Query(query, args...)}
// }

// func (q MockQuery) Iter() gocql_interface.IterInterface {
// 	return &MockIter{q.Query.Iter()}
// }

// func (q MockQuery) Exec() error {
// 	return q.Query.Exec()
// }
