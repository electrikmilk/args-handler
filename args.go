/*
 * Copyright (c) 2023 Brandon Jordan
 */

package args

import (
	"fmt"
	"os"
	"strings"
)

type argument struct {
	name         string
	short        string
	description  string
	expectsValue bool
}

// Args is a map of the args that were passed after the
// first arg with dash prefixes (e.g. -- or -) trimmed.
// A value is set for a member of Args if an arg is
// proceeded with an equality operator (e.g. --arg=value).
var Args map[string]string

var registered []argument

// CustomUsage allows you to add custom usage details.
// The value of CustomUsage is printed in between the
// name of the binary and the flags in the usage message.
var CustomUsage string

func init() {
	Args = make(map[string]string)
	if len(os.Args) <= 1 {
		return
	}
	for i, a := range os.Args {
		if i == 0 {
			continue
		}
		a = strings.TrimPrefix(a, "--")
		a = strings.TrimPrefix(a, "-")
		if strings.Contains(a, "=") {
			var keyValue = strings.Split(a, "=")
			if len(keyValue) > 1 {
				Args[keyValue[0]] = keyValue[1]
				continue
			}
		}
		Args[a] = ""
	}
}

// PrintUsage prints a usage message based on the arguments and usage you have registered then exits.
func PrintUsage() {
	var availableFlags string
	for a, arg := range registered {
		if arg.short == "" {
			availableFlags += "--" + arg.name
		} else {
			availableFlags += "-" + arg.short
		}
		if arg.expectsValue {
			availableFlags += "="
		}
		if len(registered)-1 != a {
			availableFlags += " "
		}
	}
	fmt.Printf("USAGE: %s %s [%s]", os.Args[0], CustomUsage, availableFlags)
	fmt.Printf("\nOptions:\n")
	for _, arg := range registered {
		var short = arg.short
		var name = arg.name
		if arg.expectsValue {
			short += "="
			name += "="
		} else {
			short += " "
			name += " "
		}
		if arg.short == "" {
			fmt.Printf("\t    --%s\t%s\n", name, arg.description)
		} else {
			fmt.Printf("\t-%s --%s\t%s\n", short, name, arg.description)
		}
	}
}

// Register an argument.
func Register(name string, shorthand string, description string, expectsValue bool) {
	for _, r := range registered {
		if r.name == name {
			return
		}
	}
	registered = append(registered, argument{
		name:         name,
		short:        shorthand,
		description:  description,
		expectsValue: expectsValue,
	})
}

// Using returns a boolean indicating if argument name was passed to your executable.
func Using(name string) bool {
	if len(Args) > 0 {
		if _, ok := Args[name]; ok {
			return true
		}
		for _, r := range registered {
			if r.name == name {
				if _, ok := Args[r.short]; ok {
					return true
				}
			}
		}
	}
	return false
}

// Value returns a string of the value of argument name if passed to your executable.
func Value(name string) (value string) {
	if len(Args) == 0 {
		return ""
	}
	if val, ok := Args[name]; ok {
		value = val
	}
	for _, r := range registered {
		if r.name == name {
			if val, ok := Args[r.short]; ok {
				value = val
			}
		}
	}
	return
}
