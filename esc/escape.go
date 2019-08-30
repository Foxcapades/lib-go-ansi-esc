package esc

import "fmt"

const (
	escape  = "\x1b"
  formats = "%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d"
  leader  = "["
  tailFmt = "m"
)

func FormatText(txt string, fmt, clear uint8) string {
	return EscapeFormat(fmt) + txt + EscapeFormat(clear)
}

func EscapeFormat(code uint8) string {
	return fmt.Sprintf(escape + leader + "%d" + tailFmt, code)
}

func EscapeFormats(code ...uint8) string {
	return fmt.Sprintf(mkEsc(len(code), tailFmt), genify(code)...)
}

func mkEsc(ln int, tail string) string {
	return escape + leader + formats[:ln*3-1] + tail
}

func genify(in []uint8) []interface{} {
	out := make([]interface{}, len(in))

	for i := range in {
		out[i] = in[i]
	}

	return out
}
