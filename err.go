package uu

//Err contains information about the failure
type Err string

//Error returns printable string of the failure
func (meErr Err) Error() string {
	return string(meErr)
}

const (
	ErrInvalidLineLen = Err("Invalid line length!")
	ErrInvalidDataLen = Err("Invalid data length!")
	ErrInvalidData    = Err("Invalid data!")
)
