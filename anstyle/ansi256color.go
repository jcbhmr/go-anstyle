package anstyle

import "fmt"

type ANSI256Color uint8

func ANSI256ColorFromANSI(c ANSIColor) ANSI256Color {
	return ANSI256Color(c)
}

func (c ANSI256Color) Index() uint8 {
	return uint8(c)
}

func (c ANSI256Color) IntoANSI() (ANSIColor, bool) {
	if c < 16 {
		return ANSIColor(c), true
	} else {
		return 0, false
	}
}

func (c ANSI256Color) On(background Color) Style {
	return Style{fg: optionSome(ColorANSI256(c)), bg: optionSome(background)}
}
func (c ANSI256Color) OnDefault() Style {
	return Style{fg: optionSome(ColorANSI256(c))}
}

func (c ANSI256Color) RenderFg() fmt.Stringer {
	return stringerString(fmt.Sprintf("\x1b[38;5;%dm", c))
}
func (c ANSI256Color) RenderBg() fmt.Stringer {
	return stringerString(fmt.Sprintf("\x1b[48;5;%dm", c))
}
func (c ANSI256Color) renderUnderline() fmt.Stringer {
	return stringerString(fmt.Sprintf("\x1b[58;5;%dm", c))
}
