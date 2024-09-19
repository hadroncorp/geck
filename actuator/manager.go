package actuator

import (
	"context"
	"errors"
	"os/exec"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/samber/lo"
	"go.uber.org/fx"
	"golang.org/x/sync/semaphore"

	"github.com/hadroncorp/geck/application"
	"github.com/hadroncorp/geck/internal/reflection"
)

// Manager is the Actuator coordinator. Handles heartbeat operations for health checking and holds current system
// information (Info) as well.
type Manager struct {
	Config    ConfigManager
	Actuators []Actuator
	AppConfig application.Config

	inFlightLock  sync.Mutex
	wg            sync.WaitGroup
	procSemaphore *semaphore.Weighted

	cachedInfo Info
}

// NewManagerParams Manager dependencies.
type NewManagerParams struct {
	fx.In

	Config    ConfigManager
	AppConfig application.Config
	Actuators []Actuator `group:"actuators"`
}

// NewManager allocates a new Manager instance.
func NewManager(params NewManagerParams) *Manager {
	return &Manager{
		Config:        params.Config,
		AppConfig:     params.AppConfig,
		Actuators:     params.Actuators,
		procSemaphore: semaphore.NewWeighted(lo.Max([]int64{params.Config.MaxGoroutines / 2, 4})),
	}
}

// GlobalState aggregate structure containing a global system status along registered Actuator State(s).
type GlobalState struct {
	// Status the global status of the system.
	Status Status `json:"status"`
	// Components set of State(s) from registered Actuator(s).
	Components map[string]State `json:"components"`
}

// Health returns a GlobalState aggregate. Will set global Status as StatusDown if any of registered components is down.
//
// Moreover, uses ConfigManager.MaxGoroutines value to limit the number of concurrent requests to registered Actuator
// instances as they might be remote calls to external components.
func (m *Manager) Health(ctx context.Context) (*GlobalState, error) {
	m.inFlightLock.Lock()
	defer m.inFlightLock.Unlock()
	m.wg.Add(len(m.Actuators))
	globalState := &GlobalState{
		Components: make(map[string]State),
	}
	errs := make([]error, 0, len(m.Actuators))
	writeMu := sync.Mutex{}
	for _, actuator := range m.Actuators {
		actuatorName := strings.TrimPrefix(reflection.NewTypeFullNameAny(actuator), "*")
		if err := m.procSemaphore.Acquire(ctx, 1); err != nil {
			return nil, err
		}
		go func(actuatorCopy Actuator, actuatorNameCopy string) {
			defer m.procSemaphore.Release(1)
			defer m.wg.Done()
			state, err := actuatorCopy.State(ctx)
			if err != nil {
				writeMu.Lock()
				errs = append(errs, err)
				writeMu.Unlock()
				return
			}
			writeMu.Lock()
			globalState.Components[actuatorNameCopy] = state
			writeMu.Unlock()
		}(actuator, actuatorName)
	}
	m.wg.Wait()
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	for _, componentState := range globalState.Components {
		if componentState.Status > globalState.Status {
			globalState.Status = componentState.Status
		}
	}
	return globalState, nil
}

// Info returns information about the running system (e.g. app name, environment, version, runtime information).
//
// Executes 'git' shell command if no BuildInfo.VcsVersion is not found by std debug package.
func (m *Manager) Info(ctx context.Context) Info {
	var zeroVal Info
	if m.cachedInfo != zeroVal {
		return m.cachedInfo
	}

	info := Info{
		ApplicationName: m.AppConfig.ApplicationName,
		Version:         m.AppConfig.Version,
		Environment:     m.AppConfig.Environment,
		RuntimeInfo: RuntimeInfo{
			GoVersion: runtime.Version(),
			NumCPU:    runtime.NumCPU(),
			GoOS:      runtime.GOOS,
			GoArch:    runtime.GOARCH,
		},
	}
	defer func() {
		m.cachedInfo = info
	}()
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return info
	}

	buildInfoOut := BuildInfo{}
	buildInfoOut.GoVersion = buildInfo.GoVersion
	for _, setting := range buildInfo.Settings {
		switch setting.Key {
		case "vcs":
			buildInfoOut.VcsType = setting.Value
		case "vcs.revision":
			buildInfoOut.VcsVersion = setting.Value
		case "vcs.time":
			buildInfoOut.VcsTime = setting.Value
		case "vcs.modified":
			buildInfoOut.VcsModified = setting.Value
		case "GOOS":
			buildInfoOut.GoOS = setting.Value
		case "GOARCH":
			buildInfoOut.GoArch = setting.Value
		}
	}

	if buildInfoOut.VcsVersion == "" {
		commitMetadata, err := exec.CommandContext(ctx, "git", "rev-parse", "HEAD").Output()
		if err != nil {
			return info
		}
		buildInfoOut.VcsVersion = strings.TrimSpace(string(commitMetadata))
	}
	info.BuildInfo = buildInfoOut
	return info
}
