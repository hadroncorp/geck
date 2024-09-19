package actuator

// RuntimeInfo informational structure containing runtime information (e.g. go version, host cpu number).
type RuntimeInfo struct {
	// GoVersion running Golang version.
	GoVersion string `json:"go_version"`
	// NumCPU number of CPUs of the running host.
	NumCPU int `json:"num_cpus"`
	// GoOS operating system of the running host.
	GoOS string `json:"go_os"`
	// GoArch CPU architecture of the running host.
	GoArch string `json:"go_arch"`
}

// BuildInfo informational structure containing metadata about compiled program binaries.
type BuildInfo struct {
	// GoVersion Golang version used to compile program binaries.
	GoVersion string `json:"go_version"`
	// GoOS target operating system name used to compile program binaries.
	GoOS string `json:"go_os"`
	// GoArch target CPU architecture used to compile program binaries.
	GoArch string `json:"go_arch"`
	// VcsType the version control system for the source tree where the build ran
	VcsType string `json:"vcs_type"`
	// VcsVersion the revision identifier for the current commit or checkout.
	VcsVersion string `json:"vcs_version"`
	// VcsTime the modification time associated with vcs.revision, in RFC3339 format.
	VcsTime string `json:"vcs_time"`
	// VcsModified true or false indicating whether the source tree had local modifications.
	VcsModified string `json:"vcs_modified"`
}

// Info informational structure containing several insightful information regarding the running program (e.g. app name,
// runtime and build information).
type Info struct {
	ApplicationName string      `json:"application_name"`
	Version         string      `json:"version"`
	Environment     string      `json:"environment"`
	RuntimeInfo     RuntimeInfo `json:"runtime_info"`
	BuildInfo       BuildInfo   `json:"build_info"`
}
