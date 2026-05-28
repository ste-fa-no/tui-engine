//go:build !windows

package backend

type BackendUnix struct {
	baseBackend
}

func NewBackend() Backend {
	return &BackendUnix{}
}
