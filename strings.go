package sutils

import (
	"runtime"
	"strconv"
	"strings"
)

func JoinQuoted(lines []string, quote, separator string) string {
	var joined strings.Builder
	for i, line := range lines {
		if i > 0 {
			joined.WriteString(separator)
		}
		joined.WriteString(quote)
		joined.WriteString(line)
		joined.WriteString(quote)
	}
	return joined.String()
}

func JoinQuotedInt(lines []int, quote, separator string) string {
	var joined strings.Builder
	for i, line := range lines {
		if i > 0 {
			joined.WriteString(separator)
		}
		joined.WriteString(quote)
		joined.WriteString(strconv.Itoa(line))
		joined.WriteString(quote)
	}
	return joined.String()
}

func FuncName() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	funcNameBeginning := strings.LastIndex(frame.Function, ".") + 1
	if 0 < funcNameBeginning && funcNameBeginning < len(frame.Function) {
		return ToUnderscore(frame.Function[funcNameBeginning:])
	}
	return "unknown_function"
}

func ToUnderscore(s string) string {
	const caseDifference = 'A' - 'a'
	var underscored strings.Builder
	for i, r := range s {
		if 'A' <= r && r <= 'Z' {
			if i > 0 {
				underscored.WriteRune('_')
			}
			underscored.WriteRune(r - caseDifference)
		} else {
			underscored.WriteRune(r)
		}
	}
	return underscored.String()
}

func JoinParams(first, second string) string {
	return first + "#" + second
}

func GetFirst(str string) string {
	idx := strings.Index(str, "#")
	if idx == -1 {
		return str
	}
	return str[:idx]
}

func GetSecond(str string) string {
	idx := strings.Index(str, "#")
	if idx == -1 {
		return ""
	}
	return str[idx+1:]
}
