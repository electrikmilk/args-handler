/*
 * Copyright (c) 2022 Brandon Jordan
 */

package args

import (
	"fmt"
	"os"
	"strings"
)

type argument struct {
	name        string
	short       string
	description string
}

var Args map[string]string
var registered []argument

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
		availableFlags += "-" + arg.short
		if len(registered)-1 != a {
			availableFlags += " "
		}
	}
	fmt.Printf("USAGE: %s %s [%s]", os.Args[0], CustomUsage, availableFlags)
	fmt.Printf("\nOptions:\n")
	for _, arg := range registered {
		fmt.Printf("\t-%s --%s\t%s\n", arg.short, arg.name, arg.description)
	}
	os.Exit(1)
}

// Register an argument.
func Register(name string, shorthand string, description string) {
	for _, r := range registered {
		if r.name == name {
			return
		}
	}
	registered = append(registered, argument{
		name:        name,
		short:       shorthand,
		description: description,
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
