package helper

import (
	"strconv"

	"github.com/c2h5oh/datasize"
)

func HumanSize(size string) (string, error) {
	if size == "" {
		return "", nil
	}
	s, err := strconv.Atoi(size)
	if err != nil {
		return "", err
	}
	return datasize.ByteSize(s).HumanReadable(), nil
}
