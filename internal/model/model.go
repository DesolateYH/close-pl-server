package model

import (
	"close-pl-server/internal/consts"
	"close-pl-server/internal/myerrors"
	"encoding/json"
	"github.com/DesolateYH/libary-yh-go/logger"
	"go.uber.org/zap"
)

type statsEventArgs struct {
	CpuAbsolute      float64 `json:"cpu_absolute"`
	DiskBytes        int64   `json:"disk_bytes"`
	MemoryBytes      int64   `json:"memory_bytes"`
	MemoryLimitBytes int64   `json:"memory_limit_bytes"`
	Network          struct {
		RxBytes int   `json:"rx_bytes"`
		TxBytes int64 `json:"tx_bytes"`
	} `json:"network"`
	State  string `json:"state"`
	Uptime int    `json:"uptime"`
}

type Body struct {
	Event string   `json:"event,omitempty"`
	Args  []string `json:"args,omitempty"`
}

func (b Body) getStatsEventArgs() (statsEventArgs, error) {
	if b.Event != consts.eventStats {
		logger.Get().Error("event status is not stats",
			zap.Any("body", b))
		return statsEventArgs{}, myerrors.EventIsNotStatsError
	}
	if len(b.Args) == 0 {
		logger.Get().Error("args is empty",
			zap.Any("body", b))
		return statsEventArgs{}, myerrors.ArgsIsEmptyError
	}
	var args statsEventArgs
	err := json.Unmarshal([]byte(b.Args[0]), &args)
	if err != nil {
		logger.Get().Error("fail to unmarshal statusEventArgs",
			zap.Error(err))
		return statsEventArgs{}, err
	}

	return args, nil
}
