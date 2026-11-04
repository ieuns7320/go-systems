package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: hexdump <file>")
		os.Exit(1)
	}

	path := os.Args[1]

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for i := 0; i < len(data); i += 16 {
		end := min(i+16, len(data))
		line := data[i:end]

		fmt.Printf("%08x ", i)

		for index, b := range line {
			fmt.Printf(" %02x", b)

			if index == 7 {
				fmt.Print(" ")
			}
		}

		if len(line) != 16 {
			missing := 16 - len(line)
			if len(line) > 7 {
				fmt.Print(strings.Repeat(" ", missing*3))
			} else {
				fmt.Print(strings.Repeat(" ", missing*3+1))
			}
		}

		// ascii
		fmt.Print("  |")

		for _, b := range line {
			if b >= 0x20 && b <= 0x7e {
				fmt.Printf("%c", b)
			} else {
				fmt.Print(".")
			}
		}

		fmt.Print("|")

		fmt.Println()
	}

	if len(data) != 0 {
		fmt.Printf("%08x\n", len(data))
	}
}
