package repos

type Iter interface {
	All(interface{}) error
	Close() error
	Err() error
	Next(interface{}) bool
}
