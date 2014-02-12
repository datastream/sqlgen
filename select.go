package sqlgen

import (
	"fmt"
)

// Select select ....
func Select(columns ...string) *DatabaseQuery {
	q := &DatabaseQuery{
		action:   "SELECT",
		columns:  columns,
		distinct: false,
	}
	return q
}

// Distinct "select distinct"
func (q *DatabaseQuery) Distinct() {
	q.distinct = true
}

// Limit add "limit ..."
func (q *DatabaseQuery) Limit(limit uint64) {
	q.limit = fmt.Sprintf("%d", limit)
}

// Offset add "offset ..."
func (q *DatabaseQuery) Offset(offset uint64) {
	q.offset = fmt.Sprintf("%d", offset)
}

// OrderBys add "ORDER BY"
func (q *DatabaseQuery) OrderBys(orderbys ...string) {
	q.orderBys = append(q.orderBys, orderbys...)
}
