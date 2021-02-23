package analyse

import (
	"git.iothinking.com/liudawei/LogMonitor/monitor"
	"git.iothinking.com/liudawei/LogMonitor/write"
)

type LogProcess struct {
	R chan string
	W chan string
	monitor.Reader
	write.Writer
}

func (l *LogProcess) Analyse() {
	data := <-l.R
	l.W <- data
}
