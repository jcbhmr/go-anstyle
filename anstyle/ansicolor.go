package anstyle

import "fmt"

type ANSIColor uint8

const (
	ANSIColorBlack         ANSIColor = 0
	ANSIColorRed           ANSIColor = 1
	ANSIColorGreen         ANSIColor = 2
	ANSIColorYellow        ANSIColor = 3
	ANSIColorBlue          ANSIColor = 4
	ANSIColorMagenta       ANSIColor = 5
	ANSIColorCyan          ANSIColor = 6
	ANSIColorWhite         ANSIColor = 7
	ANSIColorBrightBlack   ANSIColor = 8
	ANSIColorBrightRed     ANSIColor = 9
	ANSIColorBrightGreen   ANSIColor = 10
	ANSIColorBrightYellow  ANSIColor = 11
	ANSIColorBrightBlue    ANSIColor = 12
	ANSIColorBrightMagenta ANSIColor = 13
	ANSIColorBrightCyan    ANSIColor = 14
	ANSIColorBrightWhite   ANSIColor = 15
)

func (c ANSIColor) On(background Color) Style {
	return NewStyle().FgColor(ColorANSI(c)).BgColor(background)
}

func (c ANSIColor) OnDefault() Style {
	return NewStyle().FgColor(ColorANSI(c))
}

func (c ANSIColor) RenderFg() fmt.Stringer {
	if c >= 8 {
		return string2(fmt.Sprintf("\x1b[%dm", 90+c-8))
	} else {
		return string2(fmt.Sprintf("\x1b[%dm", 30+c))
	}
}

func (c ANSIColor) RenderBg() fmt.Stringer {
	return string2(fmt.Sprintf("\x1b[%dm", 40+c))
}

func (c ANSIColor) renderUnderline() fmt.Stringer {
	return string2(fmt.Sprintf("\x1b[%dm", 58+c))
}

func (c ANSIColor) Bright(yes bool) ANSIColor {
	if yes && c < 8 {
		return c + 8
	} else {
		return c
	}
}

func (c ANSIColor) IsBright() bool {
	return c >= 8
}
