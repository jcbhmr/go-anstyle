package anstylegit

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jcbhmr/go-anstyle/anstyle"
)

func Parse(s string) (anstyle.Style, error) {
	style := anstyle.NewStyle()
	numColors := 0
	effects := anstyle.NewEffects()
	for _, word := range strings.Fields(s) {
		switch strings.ToLower(word) {
		case "nobold":
			fallthrough
		case "no-bold":
			effects = effects.Remove(anstyle.EffectsBold)
		case "bold":
			effects = effects.Insert(anstyle.EffectsBold)
		case "nodim":
			fallthrough
		case "no-dim":
			effects = effects.Remove(anstyle.EffectsDimmed)
		case "dim":
			effects = effects.Insert(anstyle.EffectsDimmed)
		case "noul":
			fallthrough
		case "no-ul":
			effects = effects.Remove(anstyle.EffectsUnderline)
		case "ul":
			effects = effects.Insert(anstyle.EffectsUnderline)
		case "noblink":
			fallthrough
		case "no-blink":
			effects = effects.Remove(anstyle.EffectsBlink)
		case "blink":
			effects = effects.Insert(anstyle.EffectsBlink)
		case "noreverse":
			fallthrough
		case "no-reverse":
			effects = effects.Remove(anstyle.EffectsInvert)
		case "reverse":
			effects = effects.Insert(anstyle.EffectsInvert)
		case "noitalic":
			fallthrough
		case "no-italic":
			effects = effects.Remove(anstyle.EffectsItalic)
		case "italic":
			effects = effects.Insert(anstyle.EffectsItalic)
		case "nostrike":
			fallthrough
		case "no-strike":
			effects = effects.Remove(anstyle.EffectsStrikethrough)
		case "strike":
			effects = effects.Insert(anstyle.EffectsStrikethrough)
		default:
			if color, _, err := parseColor(word); err == nil {
				switch numColors {
				case 0:
					style = style.FgColor(color)
					numColors++
				case 1:
					style = style.BgColor(color)
					numColors++
				default:
					return anstyle.Style{}, ErrorExtraColor{Style: s, Word: word}
				}
			} else {
				fmt.Printf("err=%v\n", err)
				return anstyle.Style{}, ErrorUnknownWord{Style: s, Word: word}
			}
		}
	}
	style = style.Effects(effects)
	return style, nil
}

func parseColor(word string) (anstyle.Color, bool, error) {
	var color anstyle.Color
	var ok bool
	switch word {
	case "normal":
		color = nil
		ok = false
	case "-1":
		color = nil
		ok = false
	case "black":
		color = anstyle.ColorANSI(anstyle.ANSIColorBlack)
		ok = true
	case "red":
		color = anstyle.ColorANSI(anstyle.ANSIColorRed)
		ok = true
	case "green":
		color = anstyle.ColorANSI(anstyle.ANSIColorGreen)
		ok = true
	case "yellow":
		color = anstyle.ColorANSI(anstyle.ANSIColorYellow)
		ok = true
	case "blue":
		color = anstyle.ColorANSI(anstyle.ANSIColorBlue)
		ok = true
	case "magenta":
		color = anstyle.ColorANSI(anstyle.ANSIColorMagenta)
		ok = true
	case "cyan":
		color = anstyle.ColorANSI(anstyle.ANSIColorCyan)
		ok = true
	case "white":
		color = anstyle.ColorANSI(anstyle.ANSIColorWhite)
		ok = true
	default:
		if strings.HasPrefix(word, "#") && len(word) == 7 {
			r, rErr := strconv.ParseUint(word[1:3], 16, 8)
			g, gErr := strconv.ParseUint(word[3:5], 16, 8)
			b, bErr := strconv.ParseUint(word[5:7], 16, 8)
			if rErr == nil && gErr == nil && bErr == nil {
				color = anstyle.ColorRGB(anstyle.RGBColor{uint8(r), uint8(g), uint8(b)})
				ok = true
			} else {
				return nil, false, fmt.Errorf("rErr=%v gErr=%v bErr=%v", rErr, gErr, bErr)
			}
		} else if n, err := strconv.ParseUint(word, 10, 8); err == nil {
			color = anstyle.ColorANSI256(uint8(n))
			ok = true
		} else {
			return nil, false, fmt.Errorf("err=%v", err)
		}
	}
	return color, ok, nil
}
