package anstyle

import "fmt"

type Color struct {
	tag int
	a   any
}

func ColorANSI(c ANSIColor) Color {
	return Color{tag: 0, a: c}
}
func ColorANSI256(c ANSI256Color) Color {
	return Color{tag: 1, a: c}
}
func ColorRGB(c RGBColor) Color {
	return Color{tag: 2, a: c}
}

func (c Color) On(background Color) Style {
	return Style{fg: optionSome(c), bg: optionSome(background)}
}
func (c Color) OnDefault() Style {
	return Style{fg: optionSome(c)}
}

func (c Color) RenderFg() fmt.Stringer {
	switch c.tag {
	case 0:
		return c.a.(ANSIColor).RenderFg()
	case 1:
		return c.a.(ANSI256Color).RenderFg()
	case 2:
		return c.a.(RGBColor).RenderFg()
	default:
		panic("unreachable")
	}
}
func (c Color) RenderBg() fmt.Stringer {
	switch c.tag {
	case 0:
		return c.a.(ANSIColor).RenderBg()
	case 1:
		return c.a.(ANSI256Color).RenderBg()
	case 2:
		return c.a.(RGBColor).RenderBg()
	default:
		panic("unreachable")
	}
}
func (c Color) renderUnderline() fmt.Stringer {
	switch c.tag {
	case 0:
		return c.a.(ANSIColor).renderUnderline()
	case 1:
		return c.a.(ANSI256Color).renderUnderline()
	case 2:
		return c.a.(RGBColor).renderUnderline()
	default:
		panic("unreachable")
	}
}
