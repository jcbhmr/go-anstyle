package anstyle

import (
	"fmt"
	"io"
)

type Style struct {
	fg        option[Color]
	bg        option[Color]
	underline option[Color]
	effects   Effects
}

func NewStyle() Style {
	return Style{}
}

func (s Style) FgColor(color *Color) Style {
	s.fg = optionFromPtr(color)
	return s
}
func (s Style) BgColor(color *Color) Style {
	s.bg = optionFromPtr(color)
	return s
}
func (s Style) UnderlineColor(color *Color) Style {
	s.underline = optionFromPtr(color)
	return s
}
func (s Style) Effects(effects Effects) Style {
	s.effects = effects
	return s
}

func (s Style) GetFgColor() (Color, bool) {
	return s.fg.ToOk()
}
func (s Style) GetBgColor() (Color, bool) {
	return s.bg.ToOk()
}
func (s Style) GetUnderlineColor() (Color, bool) {
	return s.underline.ToOk()
}
func (s Style) GetEffects() Effects {
	return s.effects
}

func (s Style) Blink() Style {
	s.effects = s.effects.Insert(EffectsBlink)
	return s
}
func (s Style) Bold() Style {
	s.effects = s.effects.Insert(EffectsBold)
	return s
}
func (s Style) Dimmed() Style {
	s.effects = s.effects.Insert(EffectsDimmed)
	return s
}
func (s Style) Hidden() Style {
	s.effects = s.effects.Insert(EffectsHidden)
	return s
}
func (s Style) Invert() Style {
	s.effects = s.effects.Insert(EffectsInvert)
	return s
}
func (s Style) Italic() Style {
	s.effects = s.effects.Insert(EffectsItalic)
	return s
}
func (s Style) Strikethrough() Style {
	s.effects = s.effects.Insert(EffectsStrikethrough)
	return s
}
func (s Style) Underline() Style {
	s.effects = s.effects.Insert(EffectsUnderline)
	return s
}

func (s Style) Render() fmt.Stringer {
	sb := ""
	sb += s.effects.Render().String()
	if f, ok := s.fg.ToOk(); ok {
		sb += f.RenderFg().String()
	}
	if b, ok := s.bg.ToOk(); ok {
		sb += b.RenderBg().String()
	}
	if u, ok := s.underline.ToOk(); ok {
		sb += u.renderUnderline().String()
	}
	return stringerString(sb)
}
func (s Style) RenderReset() fmt.Stringer {
	if s == (Style{}) {
		return stringerString("\x1b[0m")
	} else {
		return stringerString("")
	}
}

func (s Style) WriteTo(w io.Writer) (int64, error) {
	n, err := io.WriteString(w, s.Render().String())
	return int64(n), err
}

func (s Style) WriteResetTo(w io.Writer) (int64, error) {
	n, err := io.WriteString(w, s.RenderReset().String())
	return int64(n), err
}

func (s Style) String() string {
	return s.Render().String()
}
