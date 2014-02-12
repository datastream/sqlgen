package sqlgen

// Where "where ..."
func (q *DatabaseQuery) Where(query string, args ...interface{}) {
	w := conditionPart{
		query: query,
		args:  args,
	}
	q.whereParts = append(q.whereParts, w)
}

// GroupBys add "group by ..."
func (q *DatabaseQuery) GroupBys(groupbys ...string) {
	q.groupBys = append(q.groupBys, groupbys...)
}

// Having add "having ..."
func (q *DatabaseQuery) Having(query string, args ...interface{}) {
	w := conditionPart{
		query: query,
		args:  args,
	}
	q.havingParts = append(q.havingParts, w)
}
