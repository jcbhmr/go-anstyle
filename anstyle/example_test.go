package anstyle_test

import "github.com/jcbhmr/go-anstyle/anstyle"

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func Example() {
	_ = anstyle.NewStyle().Bold()
}

func ExampleNewEffects() {
	_ = anstyle.NewEffects()
}

func ExampleEffects_IsPlain() {
	effects := anstyle.NewEffects()
	assert(effects.IsPlain())

	effects = anstyle.Effects(0)
	effects.Or(anstyle.EffectsBold, anstyle.EffectsUnderline)
	assert(!effects.IsPlain())
}

func ExampleEffects_Contains() {
	var effects anstyle.Effects
	effects.Or(anstyle.EffectsBold, anstyle.EffectsUnderline)
	assert(effects.Contains(anstyle.EffectsBold))
	
	effects = anstyle.Effects(0)
	assert(!effects.Contains(anstyle.EffectsBold))
}

func ExampleEffects_Insert() {
	effects := anstyle.NewEffects()
	effects = effects.Insert(anstyle.NewEffects())
	assert(effects.IsPlain())

	effects = anstyle.NewEffects()
	effects = effects.Insert(anstyle.EffectsBold)
	assert(effects.Contains(anstyle.EffectsBold))
}

func ExampleEffects_Remove() {
	var effects anstyle.Effects
	effects.Or(anstyle.EffectsBold, anstyle.EffectsUnderline)
	effects = effects.Remove(anstyle.EffectsBold)
	assert(!effects.Contains(anstyle.EffectsBold))
	assert(effects.Contains(anstyle.EffectsUnderline))
}