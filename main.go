package main

import (
	"git.iothinking.com/liudawei/LogMonitor/analyse"
	"git.iothinking.com/liudawei/LogMonitor/monitor"
	"git.iothinking.com/liudawei/LogMonitor/write"
	"time"
)

func main() {
	reader := &monitor.ReadFromPath{Path: "test"}
	writer := &write.WriteToInfluxDB{DbConfig: "123"}
	l := analyse.LogProcess{
		Reader: reader,
		Writer: writer,
		R:      make(chan string),
		W:      make(chan string),
	}

	go l.Read(l.R)
	go l.Analyse()
	go l.Write(l.W)

	time.Sleep(1 * time.Second)

}
