package text

import "github.com/Foxcapades/lib-go-ansi-esc/esc"

const (
	Bold = iota + 1
	Italic
	Underline
	SlowBlink
	FastBlink

	Strike   = 9
	Overline = 53
)

const (
	BoldClear = iota + 22
	ItalicClear
	UnderlineClear
	BlinkClear

	StrikeClear   = 29
	OverlineClear = 55
)

func BoldCode() string {
	return esc.EscapeFormat(Bold)
}

func BoldClearCode() string {
	return esc.EscapeFormat(BoldClear)
}

func BoldText(txt string) string {
	return esc.FormatText(txt, Bold, BoldClear)
}

func ItalicCode() string {
	return esc.EscapeFormat(Italic)
}

func ItalicClearCode() string {
	return esc.EscapeFormat(ItalicClear)
}

func ItalicText(txt string) string {
	return esc.FormatText(txt, Italic, ItalicClear)
}

func UnderlineCode() string {
	return esc.EscapeFormat(Underline)
}

func UnderlineClearCode() string {
	return esc.EscapeFormat(UnderlineClear)
}

func UnderlineText(txt string) string {
	return esc.FormatText(txt, Underline, UnderlineClear)
}

func OverlineCode() string {
	return esc.EscapeFormat(Overline)
}

func OverlineClearCode() string {
	return esc.EscapeFormat(OverlineClear)
}

func OverlineText(txt string) string {
	return esc.FormatText(txt, Overline, OverlineClear)
}

func StrikeThroughCode() string {
	return esc.EscapeFormat(Strike)
}

func StrikeThroughText(txt string) string {
	return esc.FormatText(txt, Strike, StrikeClear)
}

func SlowBlinkCode() string {
	return esc.EscapeFormat(SlowBlink)
}

func BlinkClearCode() string {
	return esc.EscapeFormat(BlinkClear)
}

func SlowBlinkText(txt string) string {
	return esc.FormatText(txt, SlowBlink, BlinkClear)
}

func FastBlinkCode() string {
	return esc.EscapeFormat(FastBlink)
}

func FastBlinkText(txt string) string {
	return esc.FormatText(txt, FastBlink, BlinkClear)
}
