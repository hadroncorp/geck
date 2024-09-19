package persistence

import (
	"time"

	"github.com/samber/lo"
)

func ConvertAuditableView(src Auditable) AuditableView {
	return AuditableView{
		CreateTime:           src.CreateTime.Format(time.RFC3339),
		CreateTimeMillis:     src.CreateTime.UnixMilli(),
		CreateBy:             lo.EmptyableToPtr(src.CreateBy),
		LastUpdateTime:       src.LastUpdateTime.Format(time.RFC3339),
		LastUpdateTimeMillis: src.LastUpdateTime.UnixMilli(),
		LastUpdateBy:         lo.EmptyableToPtr(src.LastUpdateBy),
		IsActive:             src.IsActive,
		Version:              src.Version,
	}
}
