# parsename
parsename is a Go package and tool to parse hispanic simple and compound names.

It is a fork / port of Edwood Ocasio's Python parsename https://github.com/eocasio/parsename .

## Usage

### Go package
```
import "github.com/jecolon/parsename"

n, err := parsename.New("Juan J. del Valle de la Cruz")
if err != nil { ... } // Handle parsing errors.
fmt.Println(n) // Pretty prints the name components.

// Print individual fields.
fmt.Println(n.FirstName)
fmt.Println(n.MiddleName)
fmt.Println(n.LastName)
fmt.Println(n.MaidenName)

// Reuse with different name.
n.Input = "José E. Colón"
n.Parse()
```

### Commandline tool

You can build the commandline tool in cmd with
```
go build -o parsename

Usage of ./parsename:
  -c	CSV output. First, Middle, Last, and Maiden names.
  -i	Interactive mode with prompts.
```
