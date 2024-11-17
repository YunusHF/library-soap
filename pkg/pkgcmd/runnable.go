package pkgcmd

import "context"

// Runnable defines the general interface for a runnable function which needs to be started and stopped to ensure
// graceful termination.
type Runnable interface {
	Start() error
	Stop(ctx context.Context) error
}
