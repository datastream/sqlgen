package sqlgen_test

import (
	"github.com/datastream/sqlgen"
	"testing"
)

func TestDelete(t *testing.T) {
	q := sqlgen.Delete()
	err := q.Tables("test")
	if err != nil {
		t.Fatal(err)
	}
	err = q.From("b", "c")
	if err == nil {
		t.Fatal("should not be nil")
	}
	q.Where("a = ?", 1)
	sqlstr, args, err := q.ToSQL()
	if sqlstr != "DELETE FROM test WHERE a = ?" || len(args) != 1 || err != nil {
		t.Fatal(sqlstr, args, err)
	}
}

func TestDeleteFrom(t *testing.T) {
	q := sqlgen.Delete()
	err := q.From("test")
	if err != nil {
		t.Fatal(err)
	}
	err = q.Tables("test2")
	if err == nil {
		t.Fatal("should not be nil")
	}
	err = q.From("b", "c")
	if err == nil {
		t.Fatal("should not be nil")
	}
	q.Where("a = ?", 1)
	sqlstr, args, err := q.ToSQL()
	if sqlstr != "DELETE FROM test WHERE a = ?" || len(args) != 1 || err != nil {
		t.Fatal(sqlstr, args, err)
	}
}
