package pansi

import "fmt"

func PrintErr(msg string, err error) {
	N(
		SBlue("➜"),
		" ",
		SYellow(msg),
		SRedBg(ForeWhite, err.Error()),
	)
}

func PrintWarn(msg string) {
	N(
		SBlue("➜"),
		" ",
		SYellow(msg),
	)
}

func PrintInfo(msg string) {
	N(
		SBlue("➜"),
		" ",
		SGreen(msg),
	)
}

func PrintIcon(msgs ...interface{}) {
	fmt.Print(SBlue("➜"), " ")
	fmt.Print(msgs...)
	fmt.Printf("\n")
}
