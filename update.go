package sqlgen

// Update update
func Update(table string) *DatabaseQuery {
	q := &DatabaseQuery{
		action: "UPDATE",
		tables: []string{table},
	}
	return q
}

// Set  update ... "set ..."
func (q *DatabaseQuery) Set(query string, args ...interface{}) {
	w := conditionPart{
		query: query,
		args:  args,
	}
	q.setParts = append(q.setParts, w)
}
