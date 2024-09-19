package task

import "github.com/hadroncorp/geck/data/persistence"

type Task struct {
	persistence.Auditable
	ID     string
	Name   string
	Status string
}

var _ persistence.Persistable = Task{}
