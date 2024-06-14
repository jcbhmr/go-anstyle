package anstyle

import (
	"fmt"
)

// 256-bit color support
//
// - 0..16 are [ANSIColor] palette codes
// - 16..232 map to [RGBColor] color values
// - 232.. map to [RGBColor] gray-scale values
type ANSI256Color uint8

func (c ANSI256Color) On(background Color) Style {
	return NewStyle().FgColor(ColorANSI256(c)).BgColor(background)
}

func (c ANSI256Color) OnDefault() Style {
	return NewStyle().FgColor(ColorANSI256(c))
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

func ANSI256ColorFromANSI(c ANSIColor) ANSI256Color {
	return ANSI256Color(c)
}

func (c ANSI256Color) RenderFg() fmt.Stringer {
	return string2(fmt.Sprintf("\x1b[38;5;%dm", c))
}
func (c ANSI256Color) RenderBg() fmt.Stringer {
	return string2(fmt.Sprintf("\x1b[48;5;%dm", c))
}
func (c ANSI256Color) renderUnderline() fmt.Stringer {
	return string2(fmt.Sprintf("\x1b[58;5;%dm", c))
}
