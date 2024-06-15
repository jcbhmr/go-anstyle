package anstyle

import "fmt"

// Reset terminal formatting
type Reset struct{}

// Render the ANSI code
//
// Reset also implements Display directly, so calling this method is optional.
func (r Reset) Render() fmt.Stringer {
	return string2("\x1b[0m")
}

func (r Reset) String() string {
	return r.Render().String()
}
