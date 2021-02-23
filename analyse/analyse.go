package analyse

import (
	"git.iothinking.com/liudawei/LogMonitor/monitor"
	"git.iothinking.com/liudawei/LogMonitor/write"
)

type LogProcess struct {
	r chan string
	w chan string
	monitor.Reader
	write.Writer
}

func (l *LogProcess) analyse() {
	data := <-l.r

	l.w <- data
}
