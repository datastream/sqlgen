package sqlgen

import (
	"strings"
)

type conditionPart struct {
	query string
	args  []interface{}
}

func conditionPartsToSQL(parts []conditionPart, patten string) (string, []interface{}) {
	sqls := make([]string, 0, len(parts))
	var args []interface{}
	for _, part := range parts {
		sqls = append(sqls, part.query)
		args = append(args, part.args...)
	}
	return strings.Join(sqls, patten), args
}
