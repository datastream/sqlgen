package sqlgen_test

import (
	"../sqlgen"
	"testing"
)

func TestSelect(t *testing.T) {
	q := sqlgen.Select("a")
	q.From("test")
	q.Where("a = ?", 1)
	q.Where("d = ?", "c")
	q.GroupBys("i")
	q.Having("j = ?", "k")
	q.OrderBys("a")
	q.Limit(3)
	q.Offset(7)
	q.Columns("d")
	sqlstr, args, err := q.ToSQL()
	if err != nil {
		t.Fatal(err)
	}
	if sqlstr != "SELECT a, d FROM test WHERE a = ? AND d = ? GROUP BY i HAVING j = ? ORDER BY a LIMIT 3 OFFSET 7" || len(args) != 3 {
		t.Fatal(sqlstr, args, err)
	}
}

func TestSelectALL(t *testing.T) {
	q := sqlgen.Select("")
	q.From("test")
	q.Where("a = ?", 1)
	sqlstr, args, err := q.ToSQL()
	if err == nil {
		t.Fatal(sqlstr, args, err)
	}
	q.Columns("a")
	q.Columns("d")
	sqlstr, args, err = q.ToSQL()
	if sqlstr != "SELECT a, d FROM test WHERE a = ?" || len(args) != 1 {
		t.Fatal(sqlstr, args, err)
	}
}
func TestSelectOrder(t *testing.T) {
	q := sqlgen.Select("")
	q.From("test t")
	q.Where("a = ?", 1)
        q.OrderBys("t.user_id")
	sqlstr, args, err := q.ToSQL()
	if err == nil {
		t.Fatal(sqlstr, args, err)
	}
	q.Columns("a")
	q.Columns("d")
	sqlstr, args, err = q.ToSQL()
	if sqlstr != "SELECT a, d FROM test t WHERE a = ? ORDER BY t.user_id" || len(args) != 1 {
		t.Fatal(sqlstr, args, err)
	}
}
