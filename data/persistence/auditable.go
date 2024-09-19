package persistence

import (
	"context"
	"time"

	"github.com/hadroncorp/geck/security"
)

type Auditable struct {
	CreateTime     time.Time
	CreateBy       string
	LastUpdateTime time.Time
	LastUpdateBy   string
	IsActive       bool
	Version        int64
}

var _ Persistable = Auditable{}

func NewAuditable(ctx context.Context) Auditable {
	now := time.Now().UTC()
	principal, err := security.GetPrincipalFromContext(ctx)
	var user string
	if err == nil {
		user = principal.Username()
	}
	return Auditable{
		CreateTime:     now,
		CreateBy:       user,
		LastUpdateTime: now,
		LastUpdateBy:   user,
		IsActive:       true,
		Version:        0,
	}
}

func (a Auditable) GetVersion() int64 {
	return a.Version
}

func (a *Auditable) Update(ctx context.Context) {
	a.Version++
	a.LastUpdateTime = time.Now().UTC()
	principal, err := security.GetPrincipalFromContext(ctx)
	if err == nil {
		a.LastUpdateBy = principal.Username()
	}
}
