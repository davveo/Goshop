package destroy

import (
	"Eshop/core/event"
	"Eshop/global/consts"
	"Eshop/global/variable"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func init() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM) // 监听可能的退出信号
		received := <-c                                                                           //接收信号管道中的值
		variable.ZapLog.Warn(consts.ProcessKilled, zap.String("信号值", received.String()))
		(event.CreateEventManageFactory()).FuzzyCall(variable.EventDestroyPrefix)
		close(c)
		os.Exit(1)
	}()
}
