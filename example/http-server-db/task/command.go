package task

type CreateCommand struct {
	TaskID string `json:"task_id" validate:"required"`
	Name   string `json:"name" validate:"required,lte=96"`
	Status string `json:"status" validate:"omitempty,oneof=PENDING RUNNING FAILED"`
}
