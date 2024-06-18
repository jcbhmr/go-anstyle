package anstylegit

import "fmt"

type Error interface {
	isError()
	fmt.Stringer
	error
}
type ErrorExtraColor struct {
	Style string
	Word string
}
type ErrorUnknownWord struct {
	Style string
	Word string
}

func (ErrorExtraColor) isError() {}
func (ErrorUnknownWord) isError() {}

func (e ErrorExtraColor) String() string {
	return fmt.Sprintf(`Error parsing style "%s": extra color "%s"`, e.Style, e.Word)
}
func (e ErrorUnknownWord) String() string {
	return fmt.Sprintf(`Error parsing style "%s": unknown word "%s"`, e.Style, e.Word)
}

func (e ErrorExtraColor) Error() string {
	return e.String()
}
func (e ErrorUnknownWord) Error() string {
	return e.String()
}
