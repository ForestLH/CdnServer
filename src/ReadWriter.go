package src

type ReadWriter interface {
	Write() error
	Read() ([]byte, error)
}
