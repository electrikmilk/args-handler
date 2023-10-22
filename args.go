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
	Name         string
	Short        string
	Description  string
	DefaultValue string
	Values       []string
	ExpectsValue bool
}

// Args is a map of the args that were passed after the
// first arg with dash prefixes (e.g. -- or -) trimmed.
// A value is set for a member of Args if an arg is
// proceeded with an equality operator (e.g. --arg=value).
var Args map[string]string

var registered []Argument

// CustomUsage allows you to add custom usage details.
// The value of CustomUsage is printed in between the
// Name of the binary and the flags in the usage message.
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
		var short = arg.Short
		var name = arg.Name
		if arg.ExpectsValue {
			short += "="
			name += "="
		} else {
			short += " "
			name += " "
		}

		var argumentUsage = "\t"
		if arg.Short != "" {
			argumentUsage += fmt.Sprintf(" -%s ", short)
		} else {
			argumentUsage += "    "
		}

		argumentUsage += fmt.Sprintf("\t --%s ", name)

		var argNameLength = len(arg.Name)
		if argNameLength < maxArgNameLen {
			argumentUsage += strings.Repeat(" ", maxArgNameLen-argNameLength)
		}

		argumentUsage += "\t"

		if arg.Description != "" {
			argumentUsage += fmt.Sprintf(" %s", arg.Description)
		}

		if len(arg.Values) != 0 {
			argumentUsage += " [" + strings.Join(arg.Values, ", ") + "]"
		}

		if arg.DefaultValue != "" {
			argumentUsage += fmt.Sprintf(" [default=%s]", arg.DefaultValue)
		}

		fmt.Println(argumentUsage)
	}
}

func availableFlags() (flags string) {
	for a, arg := range registered {
		if arg.Short == "" {
			flags += "--" + arg.Name
		} else {
			flags += "-" + arg.Short
		}
		if arg.ExpectsValue {
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
		var argNameLen = len(arg.Name)
		if argNameLen < max {
			continue
		}

		max = len(arg.Name)
	}

	return max
}

// Register an Argument.
func Register(arg Argument) {
	if arg.DefaultValue != "" && !arg.ExpectsValue {
		panic(fmt.Sprintf("--%s has a default value but does not expect value", arg.Name))
	}
	for _, r := range registered {
		if r.Name == arg.Name {
			panic(fmt.Sprintf("--%s is already a registred argument", arg.Name))
		}
		if arg.Short != "" && r.Short == arg.Short {
			panic(fmt.Sprintf("-%s is already a registred shorthand argument", arg.Short))
		}
	}
	registered = append(registered, arg)
}

// Using returns a boolean indicating if Argument Name was passed to your executable.
func Using(name string) bool {
	if len(Args) == 0 {
		return false
	}

	if _, ok := Args[name]; ok {
		return true
	}
	for _, r := range registered {
		if r.Name != name {
			continue
		}
		if _, ok := Args[r.Short]; ok {
			return true
		}
	}
	return false
}

// Value returns a string of the value of Argument Name if passed to your executable.
func Value(name string) string {
	if len(Args) == 0 {
		return ""
	}

	if val, ok := Args[name]; ok {
		return val
	}
	for _, r := range registered {
		if r.Name != name {
			continue
		}
		if val, ok := Args[r.Short]; ok {
			return val
		}
	}

	return ""
}
