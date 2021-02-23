package write

import (
	"fmt"
	"strings"
)

type Writer interface {
	Write(wc chan string)
}

type WriteToInfluxDB struct {
	DbConfig string
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Print(strings.ToUpper(<-wc))
}
