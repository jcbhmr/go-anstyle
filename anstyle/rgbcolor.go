package anstyle

import "fmt"

type RGBColor [3]uint8

func (c RGBColor) R() uint8 {
	return c[0]
}
func (c RGBColor) G() uint8 {
	return c[1]
}
func (c RGBColor) B() uint8 {
	return c[2]
}

func (c RGBColor) On(background Color) Style {
	return Style{fg: optionSome(ColorRGB(c)), bg: optionSome(background)}
}
func (c RGBColor) OnDefault() Style {
	return Style{fg: optionSome(ColorRGB(c))}
}

func (c RGBColor) RenderFg() fmt.Stringer {
	return stringerString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c[0], c[1], c[2]))
}
func (c RGBColor) RenderBg() fmt.Stringer {
	return stringerString(fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c[0], c[1], c[2]))
}
func (c RGBColor) renderUnderline() fmt.Stringer {
	return stringerString(fmt.Sprintf("\x1b[58;2;%d;%d;%dm", c[0], c[1], c[2]))
}
