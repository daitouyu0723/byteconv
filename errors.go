package byteconv

import "errors"

var (
	nilArray  error = errors.New("Nil byte array")
	wrongSize error = errors.New("Wrong size byte array")
	unsupType error = errors.New("Unsupported type")
)
