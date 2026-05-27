//go:build !windows

package backend

type BackendUnix struct {
	backendInit
	backendSize
	backendRestore
	backendBuffer
}

func NewBackend() Backend {
	return &BackendUnix{}
}
