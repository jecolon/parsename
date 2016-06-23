package parsename

import (
	"errors"
	"fmt"
	"strings"
)

// Name is a full human name structured into fields.
type Name struct {
	Input      string
	FirstName  string
	MiddleName string
	LastName   string
	MaidenName string
}

// String satisfies the Stringer interface thus interoperating with fmt nicely.
func (n Name) String() string {
	return fmt.Sprintf("First name: %q\nMiddle name: %q\nLast name: %q\nMaiden name: %q", n.FirstName, n.MiddleName, n.LastName, n.MaidenName)
}

// New returns a new *Name. If error isn't nil, n is invalid.
func New(s string) (n *Name, err error) {
	n = &Name{Input: s}
	err = n.Parse()
	return n, err
}

// Custom errors due to invalid input.
var (
	ErrNameInvalid = errors.New("Invalid name.")
)

// Parse processes input to produce a valid Name.
func (n *Name) Parse() (err error) {
	parts := strings.Split(strings.TrimSpace(n.Input), " ")

	// Incomplete name?
	if len(parts) < 2 || invalidLastPart(parts) {
		err = ErrNameInvalid
		return
	}

	// Start building name.
	n.FirstName = parts[0]

	// Simple name.
	if len(parts) == 2 {
		n.LastName = parts[1]
		return
	}

	// Simple full name. Trying to distinguish middle from maiden names.
	if len(parts) == 3 {
		if isAbbr(parts[1]) {
			n.MiddleName = parts[1]
			n.LastName = parts[2]
			return
		}
		if isArticle(parts[1]) || isPreposition(parts[1]) {
			n.LastName = strings.Join(parts[1:3], " ")
			return
		}
		n.LastName = parts[1]
		n.MaidenName = parts[2]
		return
	}

	// Compound full name.
	// Only initials can be middle names.
	if isInitial(parts[1]) {
		n.MiddleName = parts[1]
		parts = parts[2:]
	} else if isInitial(parts[2]) {
		n.MiddleName = strings.Join(parts[1:3], " ")
		parts = parts[3:]
	} else {
		parts = parts[1:]
	}
	// Now we should have at least a last name, maybe a maiden name and an
	// additional full second name (non initial).
	if len(parts) == 0 {
		panic("Parse: parts is empty! This shsould never happen!")
	}
	getSurnames(n, parts)
	return
}

// Private utilities start here.
func isAbbr(s string) bool {
	return strings.HasSuffix(s, ".")
}

func isInitial(s string) bool {
	return len(s) == 1 || isAbbr(s)
}

func isArticle(s string) bool {
	switch strings.ToLower(s) {
	case "la", "las", "los":
		return true
	}
	return false
}

func isPreposition(s string) bool {
	switch strings.ToLower(s) {
	case "de", "del":
		return true
	}
	return false
}

func invalidLastPart(p []string) bool {
	lastIndex := len(p) - 1
	if isAbbr(p[lastIndex]) || isArticle(p[lastIndex]) ||
		isPreposition(p[lastIndex]) || isInitial(p[lastIndex]) {
		return true
	}
	return false
}

func getSurnames(n *Name, p []string) {
	// Just the last name.
	if len(p) == 1 {
		n.LastName = p[len(p)-1]
		return
	}
	// Could be anything!
	sn, p := getSurname(p)
	if len(p) > 0 {
		n.MaidenName = sn
		n.LastName, p = getSurname(p)
		if len(p) > 0 {
			n.MiddleName = strings.Join(p, " ")
		}
	} else {
		n.LastName = sn
	}
}

func getSurname(p []string) (string, []string) {
	length := len(p)
	if length == 0 {
		panic("getSurname: p is empty! This should never happen!")
	}
	lastIndex := length - 1
	// Just a simple name. e.g. "Castro"
	if length == 1 {
		return p[lastIndex], nil
	}
	// Compound name. This is where it gets ugly.
	if isArticle(p[lastIndex-1]) || isPreposition(p[lastIndex-1]) {
		// e.g. "de la Torre"
		if length > 2 && isPreposition(p[lastIndex-2]) {
			return strings.Join(p[lastIndex-2:], " "), p[:lastIndex-2]
		}
		// e.g. "del Toro"
		return strings.Join(p[lastIndex-1:], " "), p[:lastIndex-1]
	}
	return p[lastIndex], p[:lastIndex]
}
