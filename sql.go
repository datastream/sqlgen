package sqlgen

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type DatabaseQuery struct {
	action      string
	distinct    bool
	columns     []string
	values      [][]interface{}
	tables      []string
	setParts    []conditionPart
	whereParts  []conditionPart
	havingParts []conditionPart
	groupBys    []string
	orderBys    []string
	limit       string
	offset      string
}

// From "... from ..."
func (q *DatabaseQuery) From(tables ...string) error {
	if q.action != "SELECT" && q.action != "DELETE" {
		return errors.New(q.action + " not support from function")
	}
	if q.action == "DELETE" {
		if len(tables) != 1 || len(q.tables) > 0 {
			return errors.New(q.action + " not support mulit-tables")
		}
	}
	q.tables = append(q.tables, tables...)
	return nil
}

// Tables add tables
func (q *DatabaseQuery) Tables(tables ...string) error {
	if q.action != "SELECT" && q.action != "UPDATE" && len(q.tables) > 0 {
		return errors.New(q.action + " not support mulit-tables")
	}
	q.tables = append(q.tables, tables...)
	return nil
}

// Columns add columns
func (q *DatabaseQuery) Columns(columns ...string) {
	q.columns = append(q.columns, columns...)
}

// ToSQL generate sql string
func (q *DatabaseQuery) ToSQL() (string, []interface{}, error) {
	var sqlbuf bytes.Buffer
	var args []interface{}
	var err error
	sqlbuf.WriteString(q.action + " ")
	switch q.action {
	case "SELECT":
		// select ... from ... where ...
		if q.distinct {
			sqlbuf.WriteString("DISTINCT ")
		}
		columns := strings.Join(q.columns, ", ")
		if len(columns) > 0 {
			if columns[0:2] == ", " {
				columns = columns[2:]
			}
			sqlbuf.WriteString(columns)
		} else {
			return sqlbuf.String(), args, errors.New("no columns")
		}
		if len(q.tables) > 0 {
			sqlbuf.WriteString(" FROM ")
			sqlbuf.WriteString(strings.Join(q.tables, ", "))
		} else {
			return sqlbuf.String(), args, errors.New("no tables")
		}
	case "INSERT":
		// insert into ... values ...
		if len(q.tables) == 1 {
			sqlbuf.WriteString("INTO ")
			sqlbuf.WriteString(q.tables[0])
		} else {
			return sqlbuf.String(), args, errors.New("no tables")
		}
		if len(q.columns) > 0 {
			sqlbuf.WriteString(" ( " + strings.Join(q.columns, ", ") + " ) ")
		}
		if len(q.values) > 0 {
			sqlbuf.WriteString("VALUES ")
			var sqlpart []string
			for _, v := range q.values {
				args = append(args, v...)
				var sqls []string
				for i := 0; i < len(v); i++ {
					sqls = append(sqls, "?")
				}
				sqlpart = append(sqlpart, "("+strings.Join(sqls, ",")+")")
			}
			sqlbuf.WriteString(strings.Join(sqlpart, ","))
		}
	case "UPDATE":
		// update ... set ... where ...
		if len(q.tables) > 0 {
			sqlbuf.WriteString(strings.Join(q.tables, ", "))
		} else {
			return sqlbuf.String(), args, errors.New("no tables")
		}
		sqlbuf.WriteString(" SET ")
		setSql, setArgs := conditionPartsToSQL(q.setParts, ", ")
		sqlbuf.WriteString(setSql)
		args = append(args, setArgs...)
	case "DELETE":
		// delete from ... where
		if len(q.tables) == 1 {
			sqlbuf.WriteString("FROM ")
			sqlbuf.WriteString(q.tables[0])
		} else {
			return sqlbuf.String(), args, errors.New("no tables")
		}
	default:
		return sqlbuf.String(), args, errors.New("not support")
	}
	if len(q.whereParts) > 0 {
		sqlbuf.WriteString(" WHERE ")
		whereSql, whereArgs := conditionPartsToSQL(q.whereParts, " AND ")
		sqlbuf.WriteString(whereSql)
		args = append(args, whereArgs...)
	}
	if len(q.groupBys) > 0 {
		sqlbuf.WriteString(" GROUP BY ")
		sqlbuf.WriteString(strings.Join(q.groupBys, ", "))
	}

	if len(q.havingParts) > 0 {
		sqlbuf.WriteString(" HAVING ")
		havingSql, havingArgs := conditionPartsToSQL(q.havingParts, " AND ")
		sqlbuf.WriteString(havingSql)
		args = append(args, havingArgs...)
	}
	if q.action == "SELECT" {
		if len(q.orderBys) > 0 {
			sqlbuf.WriteString(" ORDER BY ")
			valid := regexp.MustCompile("^[a-zA-Z0-9_]+((\.)[a-zA-Z0-9_]+)+?$")
			if !valid.MatchString(strings.Join(q.orderBys, "_")) {
				return sqlbuf.String(), args, fmt.Errorf("bad order by args")
			}
			sqlbuf.WriteString(strings.Join(q.orderBys, ", "))
		}
		if len(q.limit) > 0 {
			sqlbuf.WriteString(" LIMIT ")
			sqlbuf.WriteString(q.limit)
		}
		if len(q.offset) > 0 {
			sqlbuf.WriteString(" OFFSET ")
			sqlbuf.WriteString(q.offset)
		}
	}
	return sqlbuf.String(), args, err
}

func PostgresSQLFormat(query string) string {
	parts := strings.Split(query, "?")
	var sqlbuf bytes.Buffer
	for i, v := range parts {
		sqlbuf.WriteString(v)
		if len(parts) != (i + 1) {
			sqlbuf.WriteString(fmt.Sprintf("$%d", i+1))
		}
	}
	return sqlbuf.String()
}
