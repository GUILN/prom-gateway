package app

import "context"

type Application struct {
}

// Run
// Loop that runs the daemon.
func (app *Application) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		}
	}
}

// Reload.
//
// Reloads the application. To be called on system's sighup
func (app *Application) Reload() error {
	return nil
}
