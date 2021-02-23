package monitor

type Reader interface {
	Read(rc chan string)
}

type ReadFromPath struct {
	Path string
}

func (r *ReadFromPath) Read(rc chan string) {
	rc <- r.Path
}
