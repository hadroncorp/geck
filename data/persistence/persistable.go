package persistence

// Persistable a type used by Repositories to store data.
type Persistable interface {
	// GetVersion retrieves the type version delta.
	GetVersion() int64
}

// NoopPersistable the no-operation of Persistable.
type NoopPersistable struct {
	Version int64
}

var _ Persistable = NoopPersistable{}

func (n NoopPersistable) GetVersion() int64 {
	return n.Version
}
