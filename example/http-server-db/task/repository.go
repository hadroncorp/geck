package task

import "github.com/hadroncorp/geck/data/persistence"

type Repository interface {
	persistence.CrudRepository[Task, string]
}
