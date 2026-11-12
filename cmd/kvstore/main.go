package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ieuns7320/go-systems/internal/kvstore"
)

func main() {
	store := kvstore.New()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")

		ok := scanner.Scan()

		if !ok {
			break
		}

		line := scanner.Text()

		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {

		case "set":
			if len(parts) != 3 {
				fmt.Fprintln(os.Stderr, "error: usage: set <key> <value>")
				continue
			}
			store.Set(parts[1], []byte(parts[2]))
			fmt.Println("OK")

		case "get":
			if len(parts) != 2 {
				fmt.Fprintln(os.Stderr, "error: usage: get <key>")
				continue
			}
			value, err := store.Get(parts[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			fmt.Println(string(value))

		case "delete":
			if len(parts) != 2 {
				fmt.Fprintln(os.Stderr, "error: usage: delete <key>")
				continue
			}
			err := store.Delete(parts[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			fmt.Println("OK")

		case "len":
			fmt.Println(store.Len())

		case "keys":
			keys := store.Keys()
			fmt.Println(keys)

		default:
			fmt.Fprintf(os.Stderr, "error: unknown command %q\n", parts[0])
		}

	}

	err := scanner.Err()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
