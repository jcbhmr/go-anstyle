package anstyle

import (
	"fmt"
	"io"
)

type Style struct {
	fg        Color
	bg        Color
	underline Color
	effects   Effects
}

func NewStyle() Style {
	return Style{}
}

func (s Style) FgColor(color Color) Style {
	s.fg = color
	return s
}
func (s Style) BgColor(color Color) Style {
	s.bg = color
	return s
}
func (s Style) UnderlineColor(color Color) Style {
	s.underline = color
	return s
}
func (s Style) Effects(effects Effects) Style {
	s.effects = effects
	return s
}

func (s Style) GetFgColor() Color {
	return s.fg
}
func (s Style) GetBgColor() Color {
	return s.bg
}
func (s Style) GetUnderlineColor() Color {
	return s.underline
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
	if s.fg != nil {
		sb += s.fg.RenderFg().String()
	}
	if s.bg != nil {
		sb += s.bg.RenderBg().String()
	}
	if s.underline != nil {
		sb += s.underline.renderUnderline().String()
	}
	return string2(sb)
}
func (s Style) RenderReset() fmt.Stringer {
	if s.fg != nil || s.bg != nil || s.underline != nil || !s.effects.IsPlain() {
		return string2("\x1b[0m")
	} else {
		return string2("")
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
