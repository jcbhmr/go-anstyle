/*
-effect: csv of effect names to apply; optional
-layer: layer to apply effects to; optional

![](https://gist.github.com/assets/61068799/476fe491-e970-40fa-90a0-f468e65292e8)
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jcbhmr/go-anstyle/anstyle"
)

func main() {
	log.SetFlags(0)
	args, err := argsParse()
	if err != nil {
		log.Fatal(err)
	}

	for fixed := 0; fixed < 16; fixed++ {
		color, ok := anstyle.ANSI256Color(fixed).IntoANSI()
		if !ok {
			panic("unreachable")
		}
		style := style(anstyle.ColorANSI(color), args.layer, args.effects)
		printNumber(uint8(fixed), style)
		if fixed == 7 || fixed == 15 {
			fmt.Println()
		}
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
		s = anstyle.NewStyle().FgColor(color)
	case layerBg:
		s = anstyle.NewStyle().BgColor(color)
	case layerUnderline:
		s = anstyle.NewStyle().UnderlineColor(color)
	}
	s = s.Effects(s.GetEffects().Insert(effects))
	return s
}

func printNumber(fixed uint8, style anstyle.Style) (int, error) {
	return fmt.Printf("%v %X%v", style.Render(), fixed, style.RenderReset())
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

func argsParse() (args, error) {
	flag2 := flag.NewFlagSet("dump-style", flag.ContinueOnError)
	var effectFlag string
	var layerFlag string
	flag2.StringVar(&effectFlag, "effect", "", "")
	flag2.StringVar(&layerFlag, "layer", "", "")
	err := flag2.Parse(os.Args[1:])
	if err != nil {
		return args{}, err
	}
	var effects anstyle.Effects
	effectsMap := map[string]anstyle.Effects{
		"bold":             anstyle.EffectsBold,
		"dimmed":           anstyle.EffectsDimmed,
		"italic":           anstyle.EffectsItalic,
		"underline":        anstyle.EffectsUnderline,
		"double_underline": anstyle.EffectsDoubleUnderline,
		"curly_underline":  anstyle.EffectsCurlyUnderline,
		"dotted_underline": anstyle.EffectsDottedUnderline,
		"dashed_underline": anstyle.EffectsDashedUnderline,
		"blink":            anstyle.EffectsBlink,
		"invert":           anstyle.EffectsInvert,
		"hidden":           anstyle.EffectsHidden,
		"strikethrough":    anstyle.EffectsStrikethrough,
	}
	if effectFlag == "" {
		effects = anstyle.EffectsPlain
	} else {
		for _, e := range strings.Split(effectFlag, ",") {
			effects2, ok := effectsMap[e]
			if !ok {
				return args{}, fmt.Errorf("invalid effect: %s", e)
			}
			effects = effects.Insert(effects2)
		}
	}
	var layer2 layer
	layerMap := map[string]layer{
		"fg":        layerFg,
		"bg":        layerBg,
		"underline": layerUnderline,
	}
	if layerFlag == "" {
		layer2 = layerFg
	} else {
		layer3, ok := layerMap[layerFlag]
		if !ok {
			return args{}, fmt.Errorf("invalid layer: %s", layerFlag)
		}
		layer2 = layer3
	}
	return args{effects: effects, layer: layer2}, nil
}
