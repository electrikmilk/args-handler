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

var args map[string]string
var registered []argument

var customUsage string

func init() {
	args = make(map[string]string)
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
				args[keyValue[0]] = keyValue[1]
				continue
			}
		}
		args[a] = ""
	}
}

// Prints a usage message based on the arguments and usage you have registered.
func usage() {
	var availableFlags string
	for a, arg := range registered {
		availableFlags += "-" + arg.short
		if len(registered)-1 != a {
			availableFlags += " "
		}
	}
	fmt.Printf("USAGE: %s %s [%s]", os.Args[0], customUsage, availableFlags)
	fmt.Printf("\nOptions:\n")
	for _, arg := range registered {
		fmt.Printf("\t-%s --%s\t%s\n", arg.short, arg.name, arg.description)
	}
}

// Register an argument.
func registerArg(name string, shorthand string, description string) {
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

// Returns a boolean indicating if argument name was passed to your executable.
func arg(name string) bool {
	if len(args) > 0 {
		if _, ok := args[name]; ok {
			return true
		}
		for _, r := range registered {
			if r.name == name {
				if _, ok := args[r.short]; ok {
					return true
				}
			}
		}
	}
	return false
}

// Returns a string of the value of argument name if passed to your executable.
func argValue(name string) (value string) {
	if len(args) == 0 {
		return ""
	}
	if val, ok := args[name]; ok {
		value = val
	}
	for _, r := range registered {
		if r.name == name {
			if val, ok := args[r.short]; ok {
				value = val
			}
		}
	}
	return
}
