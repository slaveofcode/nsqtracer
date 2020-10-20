package pansi

import "fmt"

// Simple ANSI printing with color
// https://en.wikipedia.org/wiki/ANSI_escape_code

// f = foreground, ForeRed = foreground red color
// b = background, BackRed = background red color
const (
	ForeRed     = "31"
	BackRed     = "41"
	ForeGreen   = "32"
	BackGreen   = "42"
	ForeBlue    = "34"
	BackBlue    = "44"
	ForeYellow  = "33"
	BackYellow  = "43"
	ForeMagenta = "35"
	BackMagenta = "45"
	ForeCyan    = "36"
	BackCyan    = "46"
	ForeWhite   = "97"
	BackWhite   = "107"
	ForeGray    = "90"
	BackGray    = "100"
)

// Pansi is a struct to use at printing with color codes
type Pansi struct {
	ColorCode   string
	BgColorCode string
	Text        string
}

// PRed print text with red color
func PRed(texts ...string) {
	P(ForeRed, texts...)
}

// SRed returning text with red color
func SRed(texts ...string) string {
	return S(ForeRed, texts...)
}

// SRedBg returning text with red background and custom font color
func SRedBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackRed, texts...)
}

// PGreen print text with green color
func PGreen(texts ...string) {
	P(ForeGreen, texts...)
}

// SGreen returning text with green color
func SGreen(texts ...string) string {
	return S(ForeGreen, texts...)
}

// SGreenBg returning text with green background and custom font color
func SGreenBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackGreen, texts...)
}

// PBlue print text with blue color
func PBlue(texts ...string) {
	P(ForeBlue, texts...)
}

// SBlue returning text with blue color
func SBlue(texts ...string) string {
	return S(ForeBlue, texts...)
}

// SBlueBg returning text with blue background and custom font color
func SBlueBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackBlue, texts...)
}

// PYellow print text with yellow color
func PYellow(texts ...string) {
	P(ForeYellow, texts...)
}

// SYellow returning text with yellow color
func SYellow(texts ...string) string {
	return S(ForeYellow, texts...)
}

// SYellowBg returning text with yellow background and custom font color
func SYellowBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackYellow, texts...)
}

// PMagenta print text with magenta color
func PMagenta(texts ...string) {
	P(ForeMagenta, texts...)
}

// SMagenta returning text with magenta color
func SMagenta(texts ...string) string {
	return S(ForeMagenta, texts...)
}

// SMagentaBg returning text with white background and custom font color
func SMagentaBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackMagenta, texts...)
}

// PCyan print text with cyan color
func PCyan(texts ...string) {
	P(ForeCyan, texts...)
}

// SCyan returning text with cyan color
func SCyan(texts ...string) string {
	return S(ForeCyan, texts...)
}

// SCyanBg returning text with cyan background and custom font color
func SCyanBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackCyan, texts...)
}

// PWhite print text with white color
func PWhite(texts ...string) {
	P(ForeWhite, texts...)
}

// SWhite returning text with white color
func SWhite(texts ...string) string {
	return S(ForeWhite, texts...)
}

// SWhiteBg returning text with white background and custom font color
func SWhiteBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackWhite, texts...)
}

// PGray print text with gray color
func PGray(texts ...string) {
	P(ForeGray, texts...)
}

// SGray returning text with gray color
func SGray(texts ...string) string {
	return S(ForeGray, texts...)
}

// SGrayBg returning text with gray background and custom font color
func SGrayBg(foreColor string, texts ...string) string {
	return Sb(foreColor, BackGray, texts...)
}

// N will print text  without any color setup for default
func N(texts ...interface{}) {
	fmt.Print(texts...)
	fmt.Printf("\n")
}

// P Print text with provided color code to the console
func P(colorCode string, texts ...string) {
	var pansiTexts []Pansi
	for _, t := range texts {
		pansiTexts = append(pansiTexts, Pansi{
			ColorCode: colorCode,
			Text:      t,
		})
	}
	Print(pansiTexts)
	fmt.Printf("\n")
}

// Pb Print text with provided color and background to the console
func Pb(colorCode, bgCode string, texts ...string) {
	var pansiTexts []Pansi
	for _, t := range texts {
		pansiTexts = append(pansiTexts, Pansi{
			ColorCode:   colorCode,
			BgColorCode: bgCode,
			Text:        t,
		})
	}
	Print(pansiTexts)
	fmt.Printf("\n")
}

// S will return formatted text with color code, usually combined with P
func S(colorCode string, texts ...string) string {
	var pansiTexts []Pansi
	for _, t := range texts {
		pansiTexts = append(pansiTexts, Pansi{
			ColorCode: colorCode,
			Text:      t,
		})
	}
	return SPrint(pansiTexts)
}

// Sb will return formatted text with color and custom background color, usually combined with P
func Sb(colorCode, bgCode string, texts ...string) string {
	var pansiTexts []Pansi
	for _, t := range texts {
		pansiTexts = append(pansiTexts, Pansi{
			ColorCode:   colorCode,
			BgColorCode: bgCode,
			Text:        t,
		})
	}
	return SPrint(pansiTexts)
}

// Print will print text to the console through slices of Pansi struct
func Print(texts []Pansi) {
	for _, p := range texts {
		color := ""
		if p.ColorCode != "" {
			color = p.ColorCode
		}

		if p.BgColorCode != "" {
			if color != "" {
				color = color + ";" + p.BgColorCode
			} else {
				color = p.BgColorCode
			}
		}

		fmt.Print(string("\033["+color+"m"), string(p.Text), string("\033[0m"))
	}
}

// SPrint will return text through slices of Pansi struct
func SPrint(texts []Pansi) string {
	var t string
	for _, p := range texts {
		color := ""
		if p.ColorCode != "" {
			color = p.ColorCode
		}

		if p.BgColorCode != "" {
			if color != "" {
				color = color + ";" + p.BgColorCode
			} else {
				color = p.BgColorCode
			}
		}

		t = fmt.Sprint(string("\033["+color+"m"), string(p.Text), string("\033[0m"))
	}

	return t
}

// Clear will dissapearing printed text one line before
func Clear() {
	fmt.Println("\033[H\033[2J")
}
