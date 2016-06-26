package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jecolon/parsename"
)

type csvName struct {
	*parsename.Name
}

func (c csvName) String() string {
	return fmt.Sprintf("%q,%q,%q,%q", c.FirstName, c.MiddleName, c.LastName, c.MaidenName)
}

var interactive = flag.Bool("i", false, "Interactive mode with prompts.")
var csv = flag.Bool("c", false, "CSV output. First, Middle, Last, and Maiden names.")

func main() {
	flag.Parse()

	// Prepare Scanner.
	scanner := bufio.NewScanner(os.Stdin)

	// Doc prompt if required.
	if *interactive {
		fmt.Println("Type a name and press ENTER to process it. Type q or quit to exit.")
	}

	// Main loop.
	for {
		if *interactive {
			fmt.Print("Name: ")
		}
		// Scanning is over.
		if !scanner.Scan() {
			// Check for Scanner error.
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
			break
		}
		// Sanitize input.
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		// Check for exit request.
		if input == "q" || input == "quit" {
			break
		}
		// Go interfaces are wonderful. :)
		var n fmt.Stringer
		// Parse.
		n, err := parsename.New(input)
		// Check for parse error.
		if err != nil {
			fmt.Fprintf(os.Stderr, "parsing name %q: %v\n", input, err)
			continue
		}
		// CSV output via composition.
		if *csv {
			n = &csvName{n.(*parsename.Name)}
		}
		// Finally print it out.
		fmt.Fprintln(os.Stdout, n)
	}
}
