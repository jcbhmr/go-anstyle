package anstyle

import (
	"fmt"
	"io"
)

// ANSI Text styling
//
// You can print a Style to render the corresponding ANSI code. Using the
// alternate flag # will render the ANSI reset code, if needed. Together, this
// makes it convenient to render styles using inline format arguments.
type Style struct {
	fg        Color
	bg        Color
	underline Color
	effects   Effects
}

// No effects enabled
func NewStyle() Style {
	return Style{}
}

// Set foreground color
func (s Style) FgColor(color Color) Style {
	s.fg = color
	return s
}

// Set background color
func (s Style) BgColor(color Color) Style {
	s.bg = color
	return s
}

// Set underline color
func (s Style) UnderlineColor(color Color) Style {
	s.underline = color
	return s
}

func (s Style) Effects(effects Effects) Style {
	s.effects = effects
	return s
}

// Get the foreground color
func (s Style) GetFgColor() Color {
	return s.fg
}

// Get the background color
func (s Style) GetBgColor() Color {
	return s.bg
}

func (s Style) GetUnderlineColor() Color {
	return s.underline
}

func (s Style) GetEffects() Effects {
	return s.effects
}

// Apply blink effect
func (s Style) Blink() Style {
	s.effects = s.effects.Insert(EffectsBlink)
	return s
}

// Apply bold effect
func (s Style) Bold() Style {
	s.effects = s.effects.Insert(EffectsBold)
	return s
}

// Apply dimmed effect
func (s Style) Dimmed() Style {
	s.effects = s.effects.Insert(EffectsDimmed)
	return s
}

// Apply hidden effect
func (s Style) Hidden() Style {
	s.effects = s.effects.Insert(EffectsHidden)
	return s
}

// Apply invert effect
func (s Style) Invert() Style {
	s.effects = s.effects.Insert(EffectsInvert)
	return s
}

// Apply italic effect
func (s Style) Italic() Style {
	s.effects = s.effects.Insert(EffectsItalic)
	return s
}

// Apply strikethrough effect
func (s Style) Strikethrough() Style {
	s.effects = s.effects.Insert(EffectsStrikethrough)
	return s
}

// Apply underline effect
func (s Style) Underline() Style {
	s.effects = s.effects.Insert(EffectsUnderline)
	return s
}

// Render the ANSI code
//
// Style also implements Display directly, so calling this method is optional.
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

// Renders the relevant Reset code
//
// Unlike Reset::render, this will elide the code if there is nothing to reset.
func (s Style) RenderReset() fmt.Stringer {
	if s.fg != nil || s.bg != nil || s.underline != nil || !s.effects.IsPlain() {
		return string2("\x1b[0m")
	} else {
		return string2("")
	}
}

// Write the ANSI code
func (s Style) WriteTo(w io.Writer) (int64, error) {
	n, err := io.WriteString(w, s.Render().String())
	return int64(n), err
}

// Write the relevant Reset code
//
// Unlike Reset::render, this will elide the code if there is nothing to reset.
func (s Style) WriteResetTo(w io.Writer) (int64, error) {
	n, err := io.WriteString(w, s.RenderReset().String())
	return int64(n), err
}

func (s Style) String() string {
	return s.Render().String()
}

// Check if no styling is enabled
func (s Style) IsPlain() bool {
	return s.fg == nil && s.bg == nil && s.underline == nil && s.effects.IsPlain()
}
