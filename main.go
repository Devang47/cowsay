package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[4], lines[0], borders[5])
		ret = append(ret, s)
	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[3], lines[i], borders[3])
			ret = append(ret, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[1], lines[i], borders[2])
		ret = append(ret, s)
	}

	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return w
}

func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

func main() {
	info, _ := os.Stdin.Stat()

	var lines []string

	if info.Mode()&os.ModeCharDevice != 0 {
		var line string
    flag.StringVar(&line, "v", "", "Usage: go run main.go -v \"words to display\"")
		flag.Parse()

		if line == "" {
			fmt.Println("The command is intended to work with pipes or using the -v flag.")
			fmt.Println("Usage: fortune | gocowsay \nOR: go run main.go -v \"words to display\"")
			return
		}

		lines = append(lines, line)
	} else {
		reader := bufio.NewReader(os.Stdin)

		for {
			line, _, err := reader.ReadLine()
			if err != nil && err == io.EOF {
				break
			}
			lines = append(lines, string(line))
		}
	}

	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)
	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()
}
