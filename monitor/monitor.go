package monitor

type Reader interface {
	Read(rc chan string)
}

type ReadFromPath struct {
	path string
}

func (r *ReadFromPath) Read(rc chan string) {
	data := "test"
	rc <- data
}
