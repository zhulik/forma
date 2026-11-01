package app

import (
	"context"
	"time"

	"github.com/zhulik/pal"
)

// Run executes services through pal with shared CLI dependencies injected.
func Run(ctx context.Context, services ...pal.ServiceDef) error {
	p := pal.New(services...).
		InjectSlog().
		InitTimeout(time.Second).
		HealthCheckTimeout(time.Second).
		ShutdownTimeout(3 * time.Second)

	return p.Run(ctx)
}
