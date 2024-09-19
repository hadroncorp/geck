package task

import "github.com/hadroncorp/geck/data/persistence"

type View struct {
	TaskID string `json:"task_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	persistence.AuditableView
}
