package log

import (
	"fmt"
	"strings"
)

// var (
// 	Reset  string = "\033[0m"
// 	Red    string = "\033[31m"
// 	Green  string = "\033[32m"
// 	Yellow string = "\033[33m"
// 	Blue   string = "\033[34m"
// 	Purple string = "\033[35m"
// 	Cyan   string = "\033[36m"
// 	Gray   string = "\033[37m"
// 	White  string = "\033[97m"
// 	Bold   string = "\033[1m"
// )

type ShellFormatting int

const (
	Reset ShellFormatting = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack ShellFormatting = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack ShellFormatting = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack ShellFormatting = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack ShellFormatting = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

func ApplyFormatting(text string, formats ...ShellFormatting) string {
	colorStrings := make([]string, len(formats))
	for i, color := range formats {
		colorStrings[i] = fmt.Sprintf("\x1b[%dm", color)
	}
	colorString := strings.Join(colorStrings, "")

	return fmt.Sprintf("%s%s", colorString, text)
}

func ResetFormatting(text string) string {
	return fmt.Sprintf("%s\x1b[%dm", text, Reset)
}

// Format text using the provided directives.
// Append a reset directive, so following text will not be formatted.
// Use ApplyFormatting if this is not wanted.
func FormatText(text string, formats ...ShellFormatting) string {
	return ResetFormatting(ApplyFormatting(text, formats...))
}
