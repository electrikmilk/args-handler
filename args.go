/*
 * Copyright (c) 2023 Brandon Jordan
 */

package args

import (
	"fmt"
	"os"
	"strings"
)

type Argument struct {
	name         string
	short        string
	description  string
	defaultValue string
	values       []string
	expectsValue bool
}

// Args is a map of the args that were passed after the
// first arg with dash prefixes (e.g. -- or -) trimmed.
// A value is set for a member of Args if an arg is
// proceeded with an equality operator (e.g. --arg=value).
var Args map[string]string

var registered []Argument

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
	fmt.Printf("USAGE: %s %s [%s]", os.Args[0], CustomUsage, availableFlags())
	fmt.Printf("\nOptions:\n")
	var maxArgNameLen = argNameMaxLen()
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

		var argumentUsage = "\t"
		if arg.short != "" {
			argumentUsage += fmt.Sprintf(" -%s ", short)
		} else {
			argumentUsage += "    "
		}

		argumentUsage += fmt.Sprintf("\t --%s ", name)

		var argNameLength = len(arg.name)
		if argNameLength < maxArgNameLen {
			argumentUsage += strings.Repeat(" ", maxArgNameLen-argNameLength)
		}

		argumentUsage += "\t"

		if arg.description != "" {
			argumentUsage += fmt.Sprintf(" %s", arg.description)
		}

		if len(arg.values) != 0 {
			argumentUsage += " [" + strings.Join(arg.values, ", ") + "]"
		}

		if arg.defaultValue != "" {
			argumentUsage += fmt.Sprintf(" [default=%s]", arg.defaultValue)
		}

		fmt.Println(argumentUsage)
	}
}

func availableFlags() (flags string) {
	for a, arg := range registered {
		if arg.short == "" {
			flags += "--" + arg.name
		} else {
			flags += "-" + arg.short
		}
		if arg.expectsValue {
			flags += "="
		}
		if len(registered)-1 != a {
			flags += " "
		}
	}

	return
}

func argNameMaxLen() (max int) {
	for _, arg := range registered {
		var argNameLen = len(arg.name)
		if argNameLen < max {
			continue
		}

		max = len(arg.name)
	}

	return max
}

// Register an Argument.
func Register(arg Argument) {
	if arg.defaultValue != "" && !arg.expectsValue {
		panic(fmt.Sprintf("--%s has a default value but does not expect value", arg.name))
	}
	for _, r := range registered {
		if r.name == arg.name {
			panic(fmt.Sprintf("--%s is already a registred argument", arg.name))
		}
		if arg.short != "" && r.short == arg.short {
			panic(fmt.Sprintf("-%s is already a registred shorthand argument", arg.short))
		}
	}
	registered = append(registered, arg)
}

// Using returns a boolean indicating if Argument name was passed to your executable.
func Using(name string) bool {
	if len(Args) == 0 {
		return false
	}

	if _, ok := Args[name]; ok {
		return true
	}
	for _, r := range registered {
		if r.name != name {
			continue
		}
		if _, ok := Args[r.short]; ok {
			return true
		}
	}
	return false
}

// Value returns a string of the value of Argument name if passed to your executable.
func Value(name string) string {
	if len(Args) == 0 {
		return ""
	}

	if val, ok := Args[name]; ok {
		return val
	}
	for _, r := range registered {
		if r.name != name {
			continue
		}
		if val, ok := Args[r.short]; ok {
			return val
		}
	}

	return ""
}
