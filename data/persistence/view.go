package persistence

type AuditableView struct {
	CreateTime           string  `json:"create_time"`
	CreateTimeMillis     int64   `json:"create_time_millis"`
	CreateBy             *string `json:"create_by"`
	LastUpdateTime       string  `json:"last_update_time"`
	LastUpdateTimeMillis int64   `json:"last_update_time_millis"`
	LastUpdateBy         *string `json:"last_update_by"`
	IsActive             bool    `json:"is_active"`
	Version              int64   `json:"version"`
}
