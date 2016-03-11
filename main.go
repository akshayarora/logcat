package main

import (
	"bufio"
	_ "bytes"
	"fmt"
	"io"
	"os/exec"
	"regexp"
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

	//cmd := exec.Command("adb", `-d logcat -v time "*:S AndroidRuntime:E ZSE:*"`)
	//cmd := exec.Command("adb", "devices")
	//cmd := exec.Command("adb", "-s", "00bc1e18894429b5", "logcat", "-v", "time", "\"*:S AndroidRuntime:E ZSE:*\"")
	cmd := exec.Command("adb", "-s", "00bc1e18894429b5", "logcat", "-v", "time")
	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Errorf(err.Error())
		return
	}
	printOutput(out)
	cmd.Wait()
}

func printOutput(out io.ReadCloser) {
	in := bufio.NewScanner(out)
	//                        1 = timestamp                         2=LogLevel 3=Label 4=PID 5=Log
	re := regexp.MustCompile(`(\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3}) ([A-Z])\/(\S*)\s*\((\d*)\): (.*)`)
	for in.Scan() {
		line := in.Text()
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)
			time := matches[1]
			level := matches[2]
			label := matches[3]
			pid := matches[4]
			message := matches[5]
			fg := ForegroundWhite
			switch level {
			case "E":
				fg = ForegroundRed
			case "W":
				fg = ForegroundYellow
			}
			if "AndroidRuntime" == label || "ZSE" == label {
				fmt.Printf("%s%s %s/%s (%s):%s %s\n", fg, time, level, label, pid, message, Reset)
			}
		}
	}
}
