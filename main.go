package main

import (
	"fmt"
)

var (
	Reset string = "\x1b[0m"

	ForegroundBlack   string = "\x1b[30m"
	ForegroundRed     string = "\x1b[31m"
	ForegroundGreen   string = "\x1b[32m"
	ForegroundYellow  string = "\x1b[33m"
	ForegroundBlue    string = "\x1b[34m"
	ForegroundMagenta string = "\x1b[35m"
	ForegroundCyan    string = "\x1b[36m"
	ForegroundWhite   string = "\x1b[37m"
	ForegroundDefault string = "\x1b[39m"

	ForegroundLightBlack   string = "\x1b[90m"
	ForegroundLightRed     string = "\x1b[91m"
	ForegroundLightGreen   string = "\x1b[92m"
	ForegroundLightYellow  string = "\x1b[93m"
	ForegroundLightBlue    string = "\x1b[94m"
	ForegroundLightMagenta string = "\x1b[95m"
	ForegroundLightCyan    string = "\x1b[96m"
	ForegroundLightWhite   string = "\x1b[97m"

	BackgroundBlack   string = "\x1b[40m"
	BackgroundRed     string = "\x1b[41m"
	BackgroundGreen   string = "\x1b[42m"
	BackgroundYellow  string = "\x1b[43m"
	BackgroundBlue    string = "\x1b[44m"
	BackgroundMagenta string = "\x1b[45m"
	BackgroundCyan    string = "\x1b[46m"
	BackgroundWhite   string = "\x1b[47m"
	BackgroundDefault string = "\x1b[49m"

	BackgroundLightBlack   string = "\x1b[100m"
	BackgroundLightRed     string = "\x1b[101m"
	BackgroundLightGreen   string = "\x1b[102m"
	BackgroundLightYellow  string = "\x1b[103m"
	BackgroundLightBlue    string = "\x1b[104m"
	BackgroundLightMagenta string = "\x1b[105m"
	BackgroundLightCyan    string = "\x1b[106m"
	BackgroundLightWhite   string = "\x1b[107m"
)

func main() {
	text := "%scolored %s%stext%s\n"
	fmt.Printf(text, ForegroundRed, ForegroundBlack, BackgroundRed, Reset)
	fmt.Printf(text, ForegroundGreen, ForegroundMagenta, BackgroundGreen, Reset)
}
