package anstylegit

import (
	"testing"

	"github.com/jcbhmr/go-anstyle/anstyle"
	"github.com/stretchr/testify/assert"
)

func TestParseStyle(t *testing.T) {
	var test = func(s string, style anstyle.Style) {
		parsed, err := Parse(s)
		assert.Nil(t, err, "s=%#+v", s)
		assert.Equal(t, style, parsed, "s=%#+v", s)
	}

	test("", anstyle.NewStyle())
	test("  ", anstyle.NewStyle())
	test("normal", anstyle.NewStyle())
	test("normal normal", anstyle.NewStyle())
	test("-1 normal", anstyle.NewStyle())
	test("red", anstyle.ANSIColorRed.OnDefault())
	test("red blue", anstyle.ANSIColorRed.On(anstyle.ColorANSI(anstyle.ANSIColorBlue)))
	test("   red blue   ", anstyle.ANSIColorRed.On(anstyle.ColorANSI(anstyle.ANSIColorBlue)))
	test("red\tblue", anstyle.ANSIColorRed.On(anstyle.ColorANSI(anstyle.ANSIColorBlue)))
	test("red\n blue", anstyle.ANSIColorRed.On(anstyle.ColorANSI(anstyle.ANSIColorBlue)))
	test("red\r\n blue", anstyle.ANSIColorRed.On(anstyle.ColorANSI(anstyle.ANSIColorBlue)))
	test("blue red", anstyle.ANSIColorBlue.On(anstyle.ColorANSI(anstyle.ANSIColorRed)))
	test("yellow green", anstyle.ANSIColorYellow.On(anstyle.ColorANSI(anstyle.ANSIColorGreen)))
	test("white magenta", anstyle.ANSIColorWhite.On(anstyle.ColorANSI(anstyle.ANSIColorMagenta)))
	test("black cyan", anstyle.ANSIColorBlack.On(anstyle.ColorANSI(anstyle.ANSIColorCyan)))
	test("red normal", anstyle.ANSIColorRed.OnDefault())
	test("normal red", anstyle.NewStyle().BgColor(anstyle.ColorANSI(anstyle.ANSIColorRed)))
	test("0", anstyle.ANSI256Color(0).OnDefault())
	test("8 3", anstyle.ANSI256Color(8).On(anstyle.ColorANSI256(anstyle.ANSI256Color(3))))
	test("255", anstyle.ANSI256Color(255).OnDefault())
	test("255 -1", anstyle.ANSI256Color(255).OnDefault())
	test("#000000", anstyle.RGBColor{0, 0, 0}.OnDefault())
	test("#204060", anstyle.RGBColor{0x20, 0x40, 0x60}.OnDefault())

	test("bold cyan white", anstyle.ANSIColorCyan.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)).Bold())
	test("bold cyan nobold white", anstyle.ANSIColorCyan.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)))
	test("bold cyan reverse white nobold", anstyle.ANSIColorCyan.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)).Invert())
	test("bold cyan ul white dim", anstyle.ANSIColorCyan.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)).Bold().Underline().Dimmed())
	test("ul cyan white no-ul", anstyle.ANSIColorCyan.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)))
	test("italic cyan white", anstyle.ANSIColorCyan.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)).Italic())
	test("strike cyan white", anstyle.ANSIColorCyan.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)).Strikethrough())
	test("blink #050505 white", anstyle.RGBColor{0x05, 0x05, 0x05}.On(anstyle.ColorANSI(anstyle.ANSIColorWhite)).Blink())
}
