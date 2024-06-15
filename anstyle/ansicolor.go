package anstyle

import "fmt"

// Available 4-bit ANSI color palette codes
//
// The user’s terminal defines the meaning of the each palette code.
type ANSIColor uint8

const (
	// Black: #0 (foreground code 30, background code 40).
	ANSIColorBlack ANSIColor = 0
	// Red: #1 (foreground code 31, background code 41).
	ANSIColorRed ANSIColor = 1
	// Green: #2 (foreground code 32, background code 42).
	ANSIColorGreen ANSIColor = 2
	// Yellow: #3 (foreground code 33, background code 43).
	ANSIColorYellow ANSIColor = 3
	// Blue: #4 (foreground code 34, background code 44).
	ANSIColorBlue ANSIColor = 4
	// Magenta: #5 (foreground code 35, background code 45).
	ANSIColorMagenta ANSIColor = 5
	// Cyan: #6 (foreground code 36, background code 46).
	ANSIColorCyan ANSIColor = 6
	// White: #7 (foreground code 37, background code 47).
	ANSIColorWhite ANSIColor = 7
	// Bright black: #0 (foreground code 90, background code 100).
	ANSIColorBrightBlack ANSIColor = 8
	// Bright red: #1 (foreground code 91, background code 101).
	ANSIColorBrightRed ANSIColor = 9
	// Bright green: #2 (foreground code 92, background code 102).
	ANSIColorBrightGreen ANSIColor = 10
	// Bright yellow: #3 (foreground code 93, background code 103).
	ANSIColorBrightYellow ANSIColor = 11
	// Bright blue: #4 (foreground code 94, background code 104).
	ANSIColorBrightBlue ANSIColor = 12
	// Bright magenta: #5 (foreground code 95, background code 105).
	ANSIColorBrightMagenta ANSIColor = 13
	// Bright cyan: #6 (foreground code 96, background code 106).
	ANSIColorBrightCyan ANSIColor = 14
	// Bright white: #7 (foreground code 97, background code 107).
	ANSIColorBrightWhite ANSIColor = 15
)

// Create a Style with this as the foreground
func (c ANSIColor) On(background Color) Style {
	return NewStyle().FgColor(ColorANSI(c)).BgColor(background)
}

// Create a Style with this as the foreground
func (c ANSIColor) OnDefault() Style {
	return NewStyle().FgColor(ColorANSI(c))
}

// Render the ANSI code for a foreground color
func (c ANSIColor) RenderFg() fmt.Stringer {
	if c >= 8 {
		return string2(fmt.Sprintf("\x1b[%dm", 90+c-8))
	} else {
		return string2(fmt.Sprintf("\x1b[%dm", 30+c))
	}
}

// Render the ANSI code for a background color
func (c ANSIColor) RenderBg() fmt.Stringer {
	if c >= 8 {
		return string2(fmt.Sprintf("\x1b[%dm", 100+c-8))
	} else {
		return string2(fmt.Sprintf("\x1b[%dm", 40+c))
	}
}

func (c ANSIColor) renderUnderline() fmt.Stringer {
	return string2(fmt.Sprintf("\x1b[%dm", 58+c))
}

// Change the color to/from bright
func (c ANSIColor) Bright(yes bool) ANSIColor {
	if yes && c < 8 {
		return c + 8
	} else {
		return c
	}
}

// Report whether the color is bright
func (c ANSIColor) IsBright() bool {
	return c >= 8
}
