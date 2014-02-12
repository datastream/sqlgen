package sqlgen_test

import (
	"../sqlgen"
	"testing"
)

func TestInsert(t *testing.T) {
	q := sqlgen.Insert("test")
	err := q.Tables("test1")
	if err == nil {
		t.Fatal("not support mulit-tables")
	}
	q.Columns("a", "b", "c")
	err = q.Values([]interface{}{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	err = q.Values([]interface{}{4, 5, 6})
	sqlstr, args, err := q.ToSQL()
	if sqlstr != "INSERT INTO test ( a, b, c ) VALUES (?,?,?),(?,?,?)" || len(args) != 6 || err != nil {
		t.Fatal(sqlstr, args, err)
	}
}

func TestInsertNotMatch(t *testing.T) {
	q := sqlgen.Insert("test")
	q.Columns("a", "b", "c")
	q.Columns("d")
	err := q.Values([]interface{}{1, 2, 3})
	if err == nil {
		t.Fatal(err)
	}
}
