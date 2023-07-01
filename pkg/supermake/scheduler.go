package supermake

import "context"

type Scheduler interface {
	ScheduleTarget(ctx context.Context, execCtx ExecutorContext, target *Target) error
}

type LocalScheduler struct{}

func (s *LocalScheduler) ScheduleTarget(ctx context.Context, execCtx ExecutorContext, target *Target) error {
	return target.Run(ctx, execCtx)
}
