package repeater

import (
	"context"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

const minimumForceInterval = time.Second

type Repeater interface {
	Stop()
	Force()
}

// repeater contains logic of repeating some task.
type repeater struct {
	// Interval contains an interval between task execution.
	// Interval must be greater than zero; if not, Repeater will panic.
	interval time.Duration

	name  string
	trace trace.Driver

	// Task is a function that must be executed periodically.
	task func(context.Context) error

	runCtx context.Context
	stop   context.CancelFunc
	done   chan struct{}
	force  chan struct{}
}

type option func(r *repeater)

func WithName(name string) option {
	return func(r *repeater) {
		r.name = name
	}
}

func WithTrace(trace trace.Driver) option {
	return func(r *repeater) {
		r.trace = trace
	}
}

func WithInterval(interval time.Duration) option {
	return func(r *repeater) {
		r.interval = interval
	}
}

type event string

const (
	eventTick  = event("tick")
	eventForce = event("force")
)

// New creates and begins to execute task periodically.
func New(
	ctx context.Context,
	interval time.Duration,
	task func(ctx context.Context) (err error),
	opts ...option,
) Repeater {
	r := &repeater{
		interval: interval,
		task:     task,
		done:     make(chan struct{}),
		force:    make(chan struct{}),
	}

	r.runCtx, r.stop = context.WithCancel(context.Background())

	for _, o := range opts {
		o(r)
	}

	go r.worker(ctx, r.interval)

	return r
}

// Stop stops to execute its task.
func (r *repeater) Stop() {
	r.stop()
}

func (r *repeater) Force() {
	select {
	case r.force <- struct{}{}:
	default:
	}
}

func (r *repeater) wakeUp(ctx context.Context, e event) {
	var (
		onDone = trace.DriverOnRepeaterWakeUp(
			r.trace,
			&ctx,
			r.name,
			string(e),
		)
		cancel context.CancelFunc
		err    error
	)

	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	err = r.task(ctx)

	onDone(err)

	if err != nil {
		r.Force()
	}
}

func (r *repeater) worker(ctx context.Context, interval time.Duration) {
	defer close(r.done)

	lastForce := time.Time{}
	for {
		select {
		case <-ctx.Done():
			return
		case <-r.runCtx.Done():
			return
		case <-time.After(interval):
			r.wakeUp(ctx, eventTick)
		case <-r.force:

			// prevent overload discovery endpoint
			needSleep := minimumForceInterval - time.Since(lastForce)
			if needSleep > 0 {
				select {
				case <-ctx.Done():
					return
				case <-r.runCtx.Done():
					return
				case <-time.After(needSleep):
					// pass
				}
			}

			lastForce = time.Now()
			r.wakeUp(ctx, eventForce)
		}
	}
}
