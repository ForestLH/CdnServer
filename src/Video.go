package src

type Video struct {
	fileName string
}

func (v *Video) Read() ([]byte, error) {
	return nil, nil
}

func (v *Video) Write() error {
	return nil
}
