package sqlgen

// Delete delete
func Delete() *DatabaseQuery {
	q := &DatabaseQuery{
		action: "DELETE",
	}
	return q
}
