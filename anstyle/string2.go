package anstyle

import "fmt"

type string2 string

func (s string2) String() string {
	return string(s)
}

var _ fmt.Stringer = (*string2)(nil)
