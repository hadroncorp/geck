package task

import "github.com/hadroncorp/geck/data/persistence"

func ConvertView(src Task) View {
	return View{
		TaskID:        src.ID,
		Name:          src.Name,
		Status:        src.Status,
		AuditableView: persistence.ConvertAuditableView(src.Auditable),
	}
}
