package sqlgen

import (
	"errors"
)

// Insert insert
func Insert(table string) *DatabaseQuery {
	q := &DatabaseQuery{
		action: "INSERT",
		tables: []string{table},
	}
	return q
}

// Values insert into ... "values ..."
func (q *DatabaseQuery) Values(args []interface{}) error {
	var err error
	if len(args) != len(q.columns) {
		err = errors.New("value not match columns")
	} else {
		q.values = append(q.values, args)
	}
	return err
}
