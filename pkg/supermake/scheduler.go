package supermake

import (
	"context"
	"sync"
)

type Scheduler interface {
	ScheduleTarget(ctx context.Context, params map[string]string, execCtx ExecutorContext, target *Target) error
	TargetDone(ctx context.Context, target *Target, params map[string]string) error
	ResetTarget(ctx context.Context, target *Target, params map[string]string) error
}

type LocalScheduler struct {
	targets *sync.Map
}

func NewLocalScheduler() *LocalScheduler {
	return &LocalScheduler{
		targets: new(sync.Map),
	}
}

func (s *LocalScheduler) ScheduleTarget(ctx context.Context, params map[string]string, execCtx ExecutorContext, target *Target) error {
	lock, loaded := s.targets.LoadOrStore(target.FQNWithParameters(params), &sync.Mutex{})
	lock.(*sync.Mutex).Lock()
	defer lock.(*sync.Mutex).Unlock()

	if loaded {
		return nil
	}

	return target.Run(ctx, params, execCtx)
}

func (s *LocalScheduler) TargetDone(ctx context.Context, target *Target, params map[string]string) error {
	s.targets.Store(target.FQNWithParameters(params), &sync.Mutex{})
	return nil
}

func (s *LocalScheduler) ResetTarget(ctx context.Context, target *Target, params map[string]string) error {
	s.targets.Delete(target.FQNWithParameters(params))
	return nil
}
