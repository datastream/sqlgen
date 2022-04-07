package sqlgen_test

import (
	"github.com/datastream/sqlgen"
	"testing"
)

func TestUpdate(t *testing.T) {
	q := sqlgen.Update("test")
	q.Set("a = ?", "2")
	q.Set("d = ?", "3")
	q.Where("s = ?", 1)
	q.Where("d = 1")
	sqlstr, args, err := q.ToSQL()
	if sqlstr != "UPDATE test SET a = ?, d = ? WHERE s = ? AND d = 1" || err != nil || len(args) != 3 {
		t.Fatal(sqlstr, args, err)
	}
}

func TestUpdateWithOutWhere(t *testing.T) {
	q := sqlgen.Update("test")
	q.Set("a = ?", "2")
	q.Set("d = 1")
	sqlstr, args, err := q.ToSQL()
	if sqlstr != "UPDATE test SET a = ?, d = 1" || err != nil || len(args) != 1 {
		t.Fatal(sqlstr, args, err)
	}
}
