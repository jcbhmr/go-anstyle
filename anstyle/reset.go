package anstyle

import "fmt"

type Reset struct{}

func (r Reset) Render() fmt.Stringer {
	return stringerString("\x1b[0m")
}
