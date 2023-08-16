package supermake

import (
	"context"
	"sync"
)

type Scheduler interface {
	ScheduleTarget(ctx context.Context, execCtx ExecutorContext, target *Target) error
	TargetDone(ctx context.Context, target *Target) error
	ResetTarget(ctx context.Context, target *Target) error
}

type LocalScheduler struct {
	targets *sync.Map
}

func NewLocalScheduler() *LocalScheduler {
	return &LocalScheduler{
		targets: new(sync.Map),
	}
}

func (s *LocalScheduler) ScheduleTarget(ctx context.Context, execCtx ExecutorContext, target *Target) error {
	lock, loaded := s.targets.LoadOrStore(target.FQN(), &sync.Mutex{})
	lock.(*sync.Mutex).Lock()
	defer lock.(*sync.Mutex).Unlock()

	if loaded {
		return nil
	}

	return target.Run(ctx, execCtx)
}

func (s *LocalScheduler) TargetDone(ctx context.Context, target *Target) error {
	s.targets.Store(target.FQN(), &sync.Mutex{})
	return nil
}

func (s *LocalScheduler) ResetTarget(ctx context.Context, target *Target) error {
	s.targets.Delete(target.FQN())
	return nil
}
