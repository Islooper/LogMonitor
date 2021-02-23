package write

import "fmt"

type Writer interface {
	Writer(wc chan string)
}

type WriteToInfluxDB struct {
	dbConfig string
}

func (w *WriteToInfluxDB) Write(wc chan string) {
	fmt.Print(<-wc)

}
