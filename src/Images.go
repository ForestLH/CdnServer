package src

type Images struct {
	fileName string
}

func (image *Images) Write() error {
	return nil
}
func (image *Images) Read() ([]byte, error) {
	return nil, nil
}
