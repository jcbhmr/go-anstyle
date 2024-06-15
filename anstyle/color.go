package anstyle

import "fmt"

// Any ANSI color code scheme
type Color interface {
	isColor()
	// Create a Style with this as the foreground
	On(background Color) Style
	// Create a Style with this as the foreground
	OnDefault() Style
	// Render the ANSI code for a foreground color
	RenderFg() fmt.Stringer
	// Render the ANSI code for a background color
	RenderBg() fmt.Stringer
	renderUnderline() fmt.Stringer
}

// Available 4-bit ANSI color palette codes
//
// The user’s terminal defines the meaning of the each palette code.
type ColorANSI ANSIColor

// 256 (8-bit) color support
//
// 0..16 are [ANSIColor] palette codes
// 0..232 map to [RGBColor] color values
// 232.. map to [RGBColor] gray-scale values
type ColorANSI256 ANSI256Color

// 24-bit ANSI RGB color codes
type ColorRGB RGBColor

func (ColorANSI) isColor()    {}
func (ColorANSI256) isColor() {}
func (ColorRGB) isColor()     {}

var _ Color = (*ColorANSI)(nil)
var _ Color = (*ColorANSI256)(nil)
var _ Color = (*ColorRGB)(nil)

func (c ColorANSI) On(background Color) Style {
	return NewStyle().FgColor(c).BgColor(background)
}
func (c ColorANSI256) On(background Color) Style {
	return NewStyle().FgColor(c).BgColor(background)
}
func (c ColorRGB) On(background Color) Style {
	return NewStyle().FgColor(c).BgColor(background)
}

func (c ColorANSI) OnDefault() Style {
	return NewStyle().FgColor(c)
}
func (c ColorANSI256) OnDefault() Style {
	return NewStyle().FgColor(c)
}
func (c ColorRGB) OnDefault() Style {
	return NewStyle().FgColor(c)
}

func (c ColorANSI) RenderFg() fmt.Stringer {
	return ANSIColor(c).RenderFg()
}
func (c ColorANSI256) RenderFg() fmt.Stringer {
	return ANSI256Color(c).RenderFg()
}
func (c ColorRGB) RenderFg() fmt.Stringer {
	return RGBColor(c).RenderFg()
}

func (c ColorANSI) RenderBg() fmt.Stringer {
	return ANSIColor(c).RenderBg()
}
func (c ColorANSI256) RenderBg() fmt.Stringer {
	return ANSI256Color(c).RenderBg()
}
func (c ColorRGB) RenderBg() fmt.Stringer {
	return RGBColor(c).RenderBg()
}

func (c ColorANSI) renderUnderline() fmt.Stringer {
	return ANSIColor(c).renderUnderline()
}
func (c ColorANSI256) renderUnderline() fmt.Stringer {
	return ANSI256Color(c).renderUnderline()
}
func (c ColorRGB) renderUnderline() fmt.Stringer {
	return RGBColor(c).renderUnderline()
}
