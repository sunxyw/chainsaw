package database

type Query struct {
	conditions map[string]interface{}
}

func NewQuery() *Query {
	return &Query{
		conditions: make(map[string]interface{}),
	}
}

func (q *Query) Where(field string, value interface{}) *Query {
	q.conditions[field] = value
	return q
}

func (q *Query) Get() map[string]interface{} {
	return q.conditions
}
