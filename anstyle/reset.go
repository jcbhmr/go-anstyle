package anstyle

import "fmt"

type Reset struct{}

func (r Reset) Render() fmt.Stringer {
	return string2("\x1b[0m")
}
