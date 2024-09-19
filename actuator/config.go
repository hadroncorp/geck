package actuator

// ConfigManager configuration structure for Manager instances.
type ConfigManager struct {
	// MaxGoroutines limit the number of concurrent requests to registered Actuator
	// instances as they might be remote calls to external components.
	MaxGoroutines int64 `env:"ACTUATOR_MAX_GOROUTINES" envDefault:"4"`
}

// ConfigDiskActuator configuration structure for DiskActuator instances.
type ConfigDiskActuator struct {
	// Path mount path used by the actuator to fetch disk state.
	Path string `env:"ACTUATOR_DISK_PATH" envDefault:"/"`
	// UsedSpaceThreshold limit factor used by actuator to consider the disk as healthy/up (or not). Expressed as percentage.
	UsedSpaceThreshold float64 `env:"ACTUATOR_USED_SPACE_THRESHOLD" envDefault:"90"`
}
