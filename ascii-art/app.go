package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type flags struct {
	isJustify bool
	align     string
	font      string
	banner    map[rune][]string
	charSize  map[rune]int
}

func main() {
	args := os.Args[1:]
	errStr := errCheck(args)
	if errStr != "true" {
		fmt.Println(errStr)
		return
	}
	f := flags{isJustify: false, banner: make(map[rune][]string, 8)}
	f.align = args[2][8:]
	f.font = args[1]
	f.readFont()
}

func (f *flags) readFont() { // reading the given font and build a banner that we need
	allLetters, err := ioutil.ReadFile("./fonts/" + f.font + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	runeIndex := 31
	bannerLines := strings.Split(string(allLetters), "\n")
	for index, line := range bannerLines {
		if index%9 == 0 {
			runeIndex++
		} else {
			f.banner[rune(runeIndex)] = append(f.banner[rune(runeIndex)], line)
			f.charSize[rune(runeIndex)] = len(line)
		}
	}
}

func errCheck(args []string) string {
	if len(args) == 1 && args[0] == "--help" {
		return `Input can't be empty
type some text like: go run app.go some text
use font types: standard, shadow, thinkertoy`
	}
	if len(args) != 3 {
		return "Please, give the 3 arguments!!!"
	}
	if args[1] != "standard" && args[1] != "shadow" && args[1] != "thinkertoy" {
		return "Please, give the correct font type: standard or shadow or thinkertoy!!!"
	}
	if args[2] != "--align=center" && args[2] != "--align=left" && args[2] != "--align=right" && args[2] != "--align=justify" {
		return "Please, give the correct align type!!!"
	}
	return "true"
}

//TerminalWidth returns width of the terminal
func TerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdout
	out, _ := cmd.Output()
	return strconv.Atoi(string(out[3 : len(out)-1]))
}
