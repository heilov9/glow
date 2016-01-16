package driver

import (
	"time"

	"github.com/chrislusf/glow/util"
)

type ExecutorStatus struct {
	InputChannelStatuses []*util.ChannelStatus
	OutputChannelStatus  *util.ChannelStatus
	ReadyTime            time.Time
	StartTime            time.Time
	StopTime             time.Time
}

func (s *ExecutorStatus) Closed() bool {
	return !s.StopTime.IsZero()
}

func (s *ExecutorStatus) TimeTaken() time.Duration {
	if s.Closed() {
		return s.StopTime.Sub(s.ReadyTime)
	}
	return time.Now().Sub(s.ReadyTime)
}
