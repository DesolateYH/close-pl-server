package processor

import (
	"close-pl-server/internal/consts"
	"close-pl-server/internal/model"
	"close-pl-server/internal/myerrors"
	"errors"
	"fmt"
	"github.com/DesolateYH/libary-yh-go/logger"
	"go.uber.org/zap"
	"time"
)

const loopTime = time.Second * 10

// {"event":"stats","args":["{\"cpu_absolute\":305.601,\"disk_bytes\":2675046943,\"memory_bytes\":20518100992,\"memory_limit_bytes\":25804800000,\"network\":{\"rx_bytes\":704862335,\"tx_bytes\":3384315373},\"state\":\"running\",\"uptime\":24150801}"]}
// {"event":"console output","args":["\u003e Broadcast 3"]}

func (p *processor) Run() {
	lastMemoryUsage := float64(0)
	logger.Get().Info("begin loop send memory")
	for {
		time.Sleep(loopTime)
		resp, err := p.connection.GetResp()
		if err != nil {
			if errors.Is(err, myerrors.TokenExpireError) {
				err := p.connection.Refresh()
				if err != nil {
					return
				}
			}
			return
		}
		if resp.Event == consts.EventStats && len(resp.Args) > 0 {
			args, err := resp.GetStatsEventArgs()
			if err != nil {
				continue
			}

			memoryUsage := float64(args.MemoryBytes) / float64(args.MemoryLimitBytes) * 100
			if memoryUsage < 95 {
				logger.Get().Info("memory usage is less than 95%", zap.Float64("memoryUsage", memoryUsage))
				continue
			}
			if memoryUsage < lastMemoryUsage {
				continue
			}

			lastMemoryUsage = memoryUsage
			_, err = p.connection.SendCommend(model.Body{
				Event: consts.EventSendCommend,
				Args: []string{
					fmt.Sprintf("Broadcast current_memory_useage_%.2f%%", memoryUsage),
				},
			})
			if err != nil {
				continue
			}

			logger.Get().Info("begin restart server")
			err = p.connection.RestartServer()
			if err != nil {
				continue
			}
			time.Sleep(time.Second * 60)
		}
	}
}
