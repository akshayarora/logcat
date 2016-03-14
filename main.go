package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
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

var filterArg string
var deviceFlag bool
var deviceSerial string

type LogLevel int

const (
	Verbose LogLevel = iota
	Debug
	Info
	Warn
	Error
	Assert
	Silent

	ALL_APPS_LABEL string = "*"
)

func init() {
	flag.StringVar(&filterArg, "filter", "", "Filters to use for adb, e.g. \"*:S AndroidRuntime:E YourApp:*\"")
	flag.BoolVar(&deviceFlag, "d", true, "Read logcat from device")
	flag.StringVar(&deviceSerial, "s", "", "Serial Number for the Device to read logcat from")
}

func main() {
	flag.Parse()
	var args []string
	if deviceFlag {
		args = append(args, "-d")
	}
	if len(deviceSerial) > 0 {
		args = append(args, []string{"-s", deviceSerial}...)
	}
	args = append(args, []string{"logcat", "-v", "time"}...)
	cmd := exec.Command("adb", args...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Errorf(err.Error())
		return
	}
	var filters map[string]LogLevel
	if len(filterArg) > 0 {
		filters = parseFilters(filterArg)
	}
	printOutput(out, filters)
	cmd.Wait()
}

func parseFilters(s string) map[string]LogLevel {
	args := strings.Split(s, " ")
	filters := make(map[string]LogLevel)
	for _, arg := range args {
		parts := strings.Split(arg, ":")
		filters[parts[0]] = parseLogLevel(parts[1])
	}
	return filters
}

func shouldLog(ll LogLevel, label string, filters map[string]LogLevel) bool {
	if filters == nil {
		return true
	}
	level, ok := filters[label]
	if !ok {
		level, ok = filters[ALL_APPS_LABEL]
		if !ok {
			return false
		}
	}
	return ll >= level
}

func printOutput(out io.ReadCloser, filters map[string]LogLevel) {
	in := bufio.NewScanner(out)
	//                        1 = timestamp                         2=LogLevel 3=Label 4=PID 5=Log
	re := regexp.MustCompile(`(\d{2}-\d{2} \d{2}:\d{2}:\d{2}.\d{3}) ([A-Z])\/(\S*)\s*\(\s*(\d*)\): (.*)`)
	for in.Scan() {
		line := in.Text()
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)
			level := parseLogLevel(matches[2])
			label := matches[3]
			if shouldLog(level, label, filters) {
				time := matches[1]
				pid := matches[4]
				message := matches[5]
				fg := colorForLevel(level)
				fmt.Printf("%s%s %s/%s (%s):%s %s\n", fg, time, level, label, pid, message, Reset)
			}
		} else {
			fmt.Printf("%s%s%s\n", ForegroundRed, line, Reset)
		}
	}
}

func colorForLevel(ll LogLevel) string {
	switch ll {
	case Verbose:
		return ForegroundDefault
	case Debug:
		return ForegroundGreen
	case Info:
		return ForegroundCyan
	case Warn:
		return ForegroundYellow
	case Error:
		return ForegroundRed
	case Assert:
		return ForegroundRed
	default:
		return ForegroundDefault
	}
}

func parseLogLevel(s string) LogLevel {
	switch s {
	case "V", "*":
		return Verbose
	case "D":
		return Debug
	case "I":
		return Info
	case "W":
		return Warn
	case "E":
		return Error
	case "A":
		return Assert
	case "S":
		return Silent
	default:
		return Silent
	}
}

func (l LogLevel) String() string {
	switch l {
	case Silent:
		return "S"
	case Verbose:
		return "V"
	case Debug:
		return "D"
	case Info:
		return "I"
	case Warn:
		return "W"
	case Error:
		return "E"
	case Assert:
		return "A"
	default:
		return "S"
	}
}
