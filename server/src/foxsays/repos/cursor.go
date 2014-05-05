package repos

import (
	"labix.org/v2/mgo"
)

type Cursor interface {
	Skip(n int) Cursor
	Limit(n int) Cursor
	Sort(fields ...string) Cursor
	Iter() Iter
}

type cursor struct {
	query *mgo.Query
}

func newCursor(query *mgo.Query) Cursor {
	return &cursor{query}
}

func (c *cursor) Skip(n int) Cursor {
	c.query = c.query.Skip(n)
	return c
}

func (c *cursor) Limit(n int) Cursor {
	c.query = c.query.Limit(n)
	return c
}

func (c *cursor) Sort(fields ...string) Cursor {
	c.query = c.query.Sort(fields...)
	return c
}

func (c cursor) Iter() Iter {
	return c.query.Iter()
}
