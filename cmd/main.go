package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"local/parsename"
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

	// Need this mess for Windows console. :-/
	var stdin io.Reader = os.Stdin
	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr
	if runtime.GOOS == "windows" {
		stdin = transform.NewReader(os.Stdin, charmap.CodePage850.NewDecoder())
		stdout = transform.NewWriter(os.Stdout, charmap.CodePage850.NewEncoder())
		stderr = transform.NewWriter(os.Stderr, charmap.CodePage850.NewEncoder())
	}

	// Prepare Scanner.
	scanner := bufio.NewScanner(stdin)

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
		n, err := parsename.New(scanner.Text())
		// Check for parse error.
		if err != nil {
			fmt.Fprintln(stderr, "parsing name:", err)
			continue
		}
		// CSV output via composition.
		if *csv {
			n = &csvName{n.(*parsename.Name)}
		}
		// Finally print it out.
		fmt.Fprintln(stdout, n)
	}
}
