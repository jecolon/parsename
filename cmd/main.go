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
	if *interactive {
		fmt.Println("Type a name and press ENTER to process it. Type q or quit to exit.")
		fmt.Print("Name: ")
	}
	var stdin io.Reader = os.Stdin
	var stdout io.Writer = os.Stdout
	var stderr io.Writer = os.Stderr
	if runtime.GOOS == "windows" {
		stdin = transform.NewReader(os.Stdin, charmap.CodePage850.NewDecoder())
		stdout = transform.NewWriter(os.Stdout, charmap.CodePage850.NewEncoder())
		stderr = transform.NewWriter(os.Stderr, charmap.CodePage850.NewEncoder())
	}
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		input := strings.ToLower(scanner.Text())
		if input == "q" || input == "quit" {
			break
		}
		var n fmt.Stringer
		n, err := parsename.New(scanner.Text())
		if err != nil {
			fmt.Fprintln(stderr, "parsing name:", err)
			if *interactive {
				fmt.Print("Name: ")
			}
			continue
		}
		if *csv {
			n = &csvName{n.(*parsename.Name)}
		}
		fmt.Fprintln(stdout, n)
		if *interactive {
			fmt.Print("Name: ")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
