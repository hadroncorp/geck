package actuator

import (
	"context"

	"github.com/shirou/gopsutil/v4/disk"
)

// DiskActuator is the Actuator implementation for running host disks.
type DiskActuator struct {
	Config ConfigDiskActuator
}

var _ Actuator = (*DiskActuator)(nil)

// NewDiskActuator allocates a DiskActuator instance.
func NewDiskActuator(cfg ConfigDiskActuator) DiskActuator {
	return DiskActuator{
		Config: cfg,
	}
}

// State returns the current state of the target component. Returns error if communication with component
// has failed (not the same as State.Status).
func (d DiskActuator) State(ctx context.Context) (State, error) {
	usage, err := disk.UsageWithContext(ctx, d.Config.Path)
	if err != nil {
		return State{
			Status:      StatusDown,
			Description: err.Error(),
		}, nil
	}

	status := StatusUp
	if usage.UsedPercent >= d.Config.UsedSpaceThreshold {
		status = StatusDown
	}
	return State{
		Status: status,
		Details: map[string]any{
			"total":               usage.Total,
			"free":                usage.Free,
			"used":                usage.Used,
			"used_percent":        usage.UsedPercent,
			"inodes_total":        usage.InodesTotal,
			"inodes_free":         usage.InodesFree,
			"inodes_used":         usage.InodesUsed,
			"inodes_used_percent": usage.InodesUsedPercent,
		},
	}, nil
}
