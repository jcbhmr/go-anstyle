package anstylesvg

import "github.com/jcbhmr/go-anstyle/anstyle"

type Term struct {
	palette    *Palette
	fgColor    anstyle.Color
	bgColor    anstyle.Color
	background bool
	fontFamily string
	minWidthPx uint
	paddingPx  uint
}

var fgColor = anstyle.ColorANSI(anstyle.ANSIColorWhite)
var bgColor = anstyle.ColorANSI(anstyle.ANSIColorBlack)

func NewTerm() *Term {
	vga2 := VGA
	fgColor2 := fgColor
	bgColor2 := bgColor
	return &Term{
		palette:    &vga2,
		fgColor:    fgColor2,
		bgColor:    bgColor2,
		background: true,
		fontFamily: "SFMono-Regular, Consolas, Liberation Mono, Menlo, monospace",
		minWidthPx: 720,
		paddingPx:  10,
	}
}

func (t *Term) Palette(palette *Palette) *Term {
	t.palette = palette
	return t
}

func (t *Term) FgColor(color anstyle.Color) *Term {
	t.fgColor = color
	return t
}

func (t *Term) BgColor(color anstyle.Color) *Term {
	t.bgColor = color
	return t
}

func (t *Term) Background(yes bool) *Term {
	t.background = yes
	return t
}

func (t *Term) MinWidthPx(px uint) *Term {
	t.minWidthPx = px
	return t
}

func (t *Term) RenderSVG(ansi string) string {
	return ""
}
