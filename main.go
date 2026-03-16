package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal()
	}
	stringz := os.Args[1]
	font := os.Args[2]

	lines, err := os.ReadFile(font)
	if err != nil {
		log.Fatal(err)
	}

	splitLines := []string{}
	if font == "thinkertoy.txt" {
		splitLines = strings.Split(string(lines), "\r\n")

	} else {
		splitLines = strings.Split(string(lines), "\n")
	}
	for i := 0; i <= 8; i++ {
		for _, r := range stringz {

			start := (int(r) - 32) * 9
			fmt.Print(splitLines[start+i])
		}
		fmt.Println()
	}

}

// func asciiFont(text string) string {
// 	lines, err := os.ReadFile("standard.txt")
// 	if err == nil {
// 		log.Fatal(err)
// 	}
// 	splitLines := strings.Split(string(lines), "\n")

// }
