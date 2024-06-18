package anstylesvg

import "github.com/jcbhmr/go-anstyle/anstyle"

type Palette [16]anstyle.RGBColor

func (p *Palette) Get(color anstyle.ANSIColor) anstyle.RGBColor {
	color2 := anstyle.ANSI256ColorFromANSI(color)
	index := color2.Index()
	return p[index]
}
