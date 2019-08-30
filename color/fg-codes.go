package color

import "github.com/Foxcapades/lib-go-ansi-esc/esc"

const (
	FgBlack = iota + 30
	FgDarkRed
	FgDarkGreen
	FgYellow
	FgDarkBlue
	FgDarkMagenta
	FgDarkCyan
	FgLightGray
	FgLightGrey = FgLightGray
)

const (
	FgDarkGray = iota + 90
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgWhite
	FgDarkGrey = FgDarkGray
)

const (
	FgCustom  = 38
	FgDefault = 39
)

func FgBlackCode() string {
	return esc.EscapeFormat(FgBlack)
}

func FgDarkRedCode() string {
	return esc.EscapeFormat(FgDarkRed)
}

func FgDarkGreenCode() string {
	return esc.EscapeFormat(FgDarkGreen)
}

func FgYellowCode() string {
	return esc.EscapeFormat(FgYellow)
}

func FgDarkBlueCode() string {
	return esc.EscapeFormat(FgDarkBlue)
}

func FgDarkMagentaCode() string {
	return esc.EscapeFormat(FgDarkMagenta)
}

func FgDarkCyanCode() string {
	return esc.EscapeFormat(FgDarkCyan)
}

func FgLightGrayCode() string {
	return esc.EscapeFormat(FgLightGray)
}

func FgLightGreyCode() string {
	return esc.EscapeFormat(FgLightGrey)
}

func FgDarkGrayCode() string {
	return esc.EscapeFormat(FgDarkGray)
}

func FgLightRedCode() string {
	return esc.EscapeFormat(FgLightRed)
}

func FgLightGreenCode() string {
	return esc.EscapeFormat(FgLightGreen)
}

func FgLightYellowCode() string {
	return esc.EscapeFormat(FgLightYellow)
}

func FgLightBlueCode() string {
	return esc.EscapeFormat(FgLightBlue)
}

func FgLightMagentaCode() string {
	return esc.EscapeFormat(FgLightMagenta)
}

func FgLightCyanCode() string {
	return esc.EscapeFormat(FgLightCyan)
}

func FgWhiteCode() string {
	return esc.EscapeFormat(FgWhite)
}

func FgDarkGreyCode() string {
	return esc.EscapeFormat(FgDarkGrey)
}

func FgBlackText(txt string) string {
	return esc.FormatText(txt, FgBlack, FgDefault)
}

func FgDarkRedText(txt string) string {
	return esc.FormatText(txt, FgDarkRed, FgDefault)
}

func FgDarkGreenText(txt string) string {
	return esc.FormatText(txt, FgDarkGreen, FgDefault)
}

func FgYellowText(txt string) string {
	return esc.FormatText(txt, FgYellow, FgDefault)
}

func FgDarkBlueText(txt string) string {
	return esc.FormatText(txt, FgDarkBlue, FgDefault)
}

func FgDarkMagentaText(txt string) string {
	return esc.FormatText(txt, FgDarkMagenta, FgDefault)
}

func FgDarkCyanText(txt string) string {
	return esc.FormatText(txt, FgDarkCyan, FgDefault)
}

func FgLightGrayText(txt string) string {
	return esc.FormatText(txt, FgLightGray, FgDefault)
}

func FgLightGreyText(txt string) string {
	return esc.FormatText(txt, FgLightGrey, FgDefault)
}

func FgDarkGrayText(txt string) string {
	return esc.FormatText(txt, FgDarkGray, FgDefault)
}

func FgLightRedText(txt string) string {
	return esc.FormatText(txt, FgLightRed, FgDefault)
}

func FgLightGreenText(txt string) string {
	return esc.FormatText(txt, FgLightGreen, FgDefault)
}

func FgLightYellowText(txt string) string {
	return esc.FormatText(txt, FgLightYellow, FgDefault)
}

func FgLightBlueText(txt string) string {
	return esc.FormatText(txt, FgLightBlue, FgDefault)
}

func FgLightMagentaText(txt string) string {
	return esc.FormatText(txt, FgLightMagenta, FgDefault)
}

func FgLightCyanText(txt string) string {
	return esc.FormatText(txt, FgLightCyan, FgDefault)
}

func FgWhiteText(txt string) string {
	return esc.FormatText(txt, FgWhite, FgDefault)
}

func FgDarkGreyText(txt string) string {
	return esc.FormatText(txt, FgDarkGrey, FgDefault)
}

func FgDefaultCode() string {
	return esc.EscapeFormat(FgDefault)
}

func FgCustom8BitCode(color uint8) string {
	return esc.EscapeFormats(FgCustom, 5, color)
}

func FgCustom8BitText(txt string, color uint8) string {
	return FgCustom8BitCode(color) + txt + FgDefaultCode()
}

func FgCustom24BitCode(r, g, b uint8) string {
	return esc.EscapeFormats(FgCustom, 2, r, g, b)
}

func FgCustom24BitText(txt string, r, g, b uint8) string {
	return FgCustom24BitCode(r, g, b) + txt + FgDefaultCode()
}