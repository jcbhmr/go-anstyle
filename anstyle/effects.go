package anstyle

import "fmt"

type Effects uint16

const EffectsPlain Effects = 0
const EffectsBold Effects = 1 << 0
const EffectsDimmed Effects = 1 << 1
const EffectsItalic Effects = 1 << 2
const EffectsUnderline Effects = 1 << 3
const EffectsDoubleUnderline Effects = 1 << 4
const EffectsCurlyUnderline Effects = 1 << 5
const EffectsDottedUnderline Effects = 1 << 6
const EffectsDashedUnderline Effects = 1 << 7
const EffectsBlink Effects = 1 << 8
const EffectsInvert Effects = 1 << 9
const EffectsHidden Effects = 1 << 10
const EffectsStrikethrough Effects = 1 << 11

func NewEffects() Effects {
	return EffectsPlain
}

func (e Effects) IsPlain() bool {
	return e == EffectsPlain
}

func (e Effects) Contains(other Effects) bool {
	return e&other == other
}

func (e Effects) Insert(other Effects) Effects {
	return e | other
}

func (e Effects) Remove(other Effects) Effects {
	return e &^ other
}

func (e Effects) Set(other Effects, enable bool) Effects {
	if enable {
		return e.Insert(other)
	} else {
		return e.Remove(other)
	}
}

func (e Effects) Render() fmt.Stringer {
	sb := ""
	index := 0
	for index < len(metadata2) {
		index2 := index
		index++
		effect := Effects(1 << index2)
		if e.Contains(effect) {
			sb += metadata2[index2].escape
		}
	}
	return string2(sb)
}

type metadata struct {
	name   string
	escape string
}

var metadata2 = []metadata{
	{"BOLD", "\x1b[1m"},
	{"DIMMED", "\x1b[2m"},
	{"ITALIC", "\x1b[3m"},
	{"UNDERLINE", "\x1b[4m"},
	{"DOUBLE_UNDERLINE", "\x1b[21m"},
	{"CURLY_UNDERLINE", "\x1b[4:3m"},
	{"DOTTED_UNDERLINE", "\x1b[4:4m"},
	{"DASHED_UNDERLINE", "\x1b[4:5m"},
	{"BLINK", "\x1b[5m"},
	{"INVERT", "\x1b[7m"},
	{"HIDDEN", "\x1b[8m"},
	{"STRIKETHROUGH", "\x1b[9m"},
}
