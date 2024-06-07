package wrapper

import (
	"fmt"
	"runtime"

	"github.com/brunoga/robomaster/support/logger"
	wrapper_internal "github.com/brunoga/robomaster/unitybridge/wrapper"
)

type UnityBridge struct {
	uw wrapper_internal.UnityBridge
}

func New() (*UnityBridge, error) {
	l := logger.New(logger.LevelTrace)

	uw := wrapper_internal.Get(l)
	if uw == nil {
		return nil, fmt.Errorf("failed to get UnityBridge")
	}

	return &UnityBridge{
		uw: uw,
	}, nil
}

func (u *UnityBridge) Create(name string, debuggable bool, logPath string) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	u.uw.Create(name, debuggable, logPath)
}

func (u *UnityBridge) Initialize() bool {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	return u.uw.Initialize()
}

func (u *UnityBridge) Uninitialize() {
	u.uw.Uninitialize()
}

func (u *UnityBridge) Destroy() {
	u.uw.Destroy()
}
