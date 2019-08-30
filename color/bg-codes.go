package color

import "github.com/Foxcapades/lib-go-ansi-esc/esc"

const (
	BgBlack = iota + 40
	BgDarkRed
	BgDarkGreen
	BgYellow
	BgDarkBlue
	BgDarkMagenta
	BgDarkCyan
	BgLightGray
	BgLightGrey = BgLightGray
)

const (
	BgDarkGray = iota + 90
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhite
	BgDarkGrey = BgDarkGray
)

const (
	BgCustom  = 48
	BgDefault = 49
)

func BgBlackCode() string {
	return esc.EscapeFormat(BgBlack)
}

func BgDarkRedCode() string {
	return esc.EscapeFormat(BgDarkRed)
}

func BgDarkGreenCode() string {
	return esc.EscapeFormat(BgDarkGreen)
}

func BgYellowCode() string {
	return esc.EscapeFormat(BgYellow)
}

func BgDarkBlueCode() string {
	return esc.EscapeFormat(BgDarkBlue)
}

func BgDarkMagentaCode() string {
	return esc.EscapeFormat(BgDarkMagenta)
}

func BgDarkCyanCode() string {
	return esc.EscapeFormat(BgDarkCyan)
}

func BgLightGrayCode() string {
	return esc.EscapeFormat(BgLightGray)
}

func BgLightGreyCode() string {
	return esc.EscapeFormat(BgLightGrey)
}

func BgDarkGrayCode() string {
	return esc.EscapeFormat(BgDarkGray)
}

func BgLightRedCode() string {
	return esc.EscapeFormat(BgLightRed)
}

func BgLightGreenCode() string {
	return esc.EscapeFormat(BgLightGreen)
}

func BgLightYellowCode() string {
	return esc.EscapeFormat(BgLightYellow)
}

func BgLightBlueCode() string {
	return esc.EscapeFormat(BgLightBlue)
}

func BgLightMagentaCode() string {
	return esc.EscapeFormat(BgLightMagenta)
}

func BgLightCyanCode() string {
	return esc.EscapeFormat(BgLightCyan)
}

func BgWhiteCode() string {
	return esc.EscapeFormat(BgWhite)
}

func BgDarkGreyCode() string {
	return esc.EscapeFormat(BgDarkGrey)
}

func BgCustom8BitCode(color uint8) string {
	return esc.EscapeFormats(BgCustom, 5, color)
}

func BgCustom24BitCode(r, g, b uint8) string {
	return esc.EscapeFormats(BgCustom, 2, r, g, b)
}

func BgDefaultCode() string {
	return esc.EscapeFormat(BgDefault)
}

func BgBlackText(txt string) string {
	return BgBlackCode() + txt + BgDefaultCode()
}

func BgDarkRedText(txt string) string {
	return BgDarkRedCode() + txt + BgDefaultCode()
}

func BgDarkGreenText(txt string) string {
	return BgDarkGreenCode() + txt + BgDefaultCode()
}

func BgYellowText(txt string) string {
	return BgYellowCode() + txt + BgDefaultCode()
}

func BgDarkBlueText(txt string) string {
	return BgDarkBlueCode() + txt + BgDefaultCode()
}

func BgDarkMagentaText(txt string) string {
	return BgDarkMagentaCode() + txt + BgDefaultCode()
}

func BgDarkCyanText(txt string) string {
	return BgDarkCyanCode() + txt + BgDefaultCode()
}

func BgLightGrayText(txt string) string {
	return BgLightGrayCode() + txt + BgDefaultCode()
}

func BgLightGreyText(txt string) string {
	return BgLightGreyCode() + txt + BgDefaultCode()
}

func BgDarkGrayText(txt string) string {
	return BgDarkGrayCode() + txt + BgDefaultCode()
}

func BgLightRedText(txt string) string {
	return BgLightRedCode() + txt + BgDefaultCode()
}

func BgLightGreenText(txt string) string {
	return BgLightGreenCode() + txt + BgDefaultCode()
}

func BgLightYellowText(txt string) string {
	return BgLightYellowCode() + txt + BgDefaultCode()
}

func BgLightBlueText(txt string) string {
	return BgLightBlueCode() + txt + BgDefaultCode()
}

func BgLightMagentaText(txt string) string {
	return BgLightMagentaCode() + txt + BgDefaultCode()
}

func BgLightCyanText(txt string) string {
	return BgLightCyanCode() + txt + BgDefaultCode()
}

func BgWhiteText(txt string) string {
	return BgWhiteCode() + txt + BgDefaultCode()
}

func BgDarkGreyText(txt string) string {
	return BgDarkGreyCode() + txt + BgDefaultCode()
}

func BgCustom8BitText(txt string, color uint8) string {
	return BgCustom8BitCode(color) + txt + BgDefaultCode()
}

func BgCustom24BitText(txt string, r, g, b uint8) string {
	return BgCustom24BitCode(r, g, b) + txt + BgDefaultCode()
}
