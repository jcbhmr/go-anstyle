package main

import (
	"fmt"
	"log"

	"github.com/jcbhmr/go-anstyle/anstyle"
)

func main() {
	log.SetFlags(0)
	args := args{}

	for fixed := 0; fixed < 16; fixed++ {
		color, ok := anstyle.ANSI256Color(fixed).IntoANSI()
		if !ok {
			panic("unreachable")
		}
		style := style(anstyle.ColorANSI(color), args.layer, args.effects)
		printNumber(uint8(fixed), style)
	}

	for fixed := 16; fixed < 232; fixed++ {
		col := (fixed - 16) % 36
		if col == 0 {
			fmt.Println()
		}
		color := anstyle.ANSI256Color(fixed)
		style := style(anstyle.ColorANSI256(color), args.layer, args.effects)
		printNumber(uint8(fixed), style)
	}

	fmt.Println()
	fmt.Println()
	for fixed := 232; fixed <= 255; fixed++ {
		color := anstyle.ANSI256Color(fixed)
		style := style(anstyle.ColorANSI256(color), args.layer, args.effects)
		printNumber(uint8(fixed), style)
	}

	fmt.Println()
}

func style(color anstyle.Color, layer layer, effects anstyle.Effects) anstyle.Style {
	var s anstyle.Style
	switch layer {
	case layerFg:
		s = anstyle.NewStyle().FgColor(&color)
	case layerBg:
		s = anstyle.NewStyle().BgColor(&color)
	case layerUnderline:
		s = anstyle.NewStyle().UnderlineColor(&color)
	}
	s = s.Effects(s.GetEffects().Insert(effects))
	return s
}

func printNumber(fixed uint8, style anstyle.Style) (int, error) {
	return fmt.Printf("%v%3d%#v", style, fixed, style)
}

type args struct {
	effects anstyle.Effects
	layer   layer
}

type layer int

const (
	layerFg layer = iota
	layerBg
	layerUnderline
)

// func argsParse() (*args, error) {
// 	res := args{}
// 	args := lexopt.ParserFromEnv()
// 	for {
// 		arg, err := args.Next()
// 		if err != nil {
// 			return nil, err
// 		}
// 		if arg == nil {
// 			break
// 		}

// 		longName, _ := arg.(lexopt.ArgLong)
// 		if string(longName) == "layer" {
// 			value, err := args.Value()
// 			if err != nil {
// 				return nil, err
// 			}
// 			layerAny, err := value.ParseWith(func(s string) (any, error) {
// 				switch s {
// 				case "fg":
// 					return layerFg, nil
// 				case "bg":
// 					return layerBg, nil
// 				case "underline":
// 					return layerUnderline, nil
// 				default:
// 					return layer(0), errors.New("expected values fg, bg, underline")
// 				}
// 			})
// 			if err != nil {
// 				return nil, err
// 			}
// 			res.layer = layerAny.(Layer)
// 		} else if string(longName) == "effect" {
// 			effects := [12]struct {
// 				A string
// 				B anstyle.Effects
// 			}{
// 				{"bold", anstyle.EffectsBold},
// 				{"dimmed", anstyle.EffectsDimmed},
// 				{"italic", anstyle.EffectsItalic},
// 				{"underline", anstyle.EffectsUnderline},
// 				{"double_underline", anstyle.EffectsDoubleUnderline},
// 				{"curly_underline", anstyle.EffectsCurlyUnderline},
// 				{"dotted_underline", anstyle.EffectsDottedUnderline},
// 				{"dashed_underline", anstyle.EffectsDashedUnderline},
// 				{"blink", anstyle.EffectsBlink},
// 				{"invert", anstyle.EffectsInvert},
// 				{"hidden", anstyle.EffectsHidden},
// 				{"strikethrough", anstyle.EffectsStrikethrough},
// 			}
// 			value, err := args.Value()
// 			if err != nil {
// 				return nil, err
// 			}
// 			effectAny, err := value.ParseWith(func(s string) (any, error) {
// 				for _, effect := range effects {
// 					if effect.A == s {
// 						return effect.B, nil
// 					}
// 				}
// 				effectKeys := make([]string, len(effects))
// 				for i, effect := range effects {
// 					effectKeys[i] = effect.A
// 				}
// 				return anstyle.Effects(0), fmt.Errorf("expected one of %s", strings.Join(effectKeys, ", "))
// 			})
// 			if err != nil {
// 				return nil, err
// 			}
// 			res.effects = res.effects.Insert(effectAny.(anstyle.Effect))
// 		} else {
// 			return nil, arg.Unexpected()
// 		}
// 	}

// 	return &res, nil
// }
