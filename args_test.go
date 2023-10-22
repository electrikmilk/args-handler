/*
 * Copyright (c) 2023 Brandon Jordan
 */

package args

import (
	"fmt"
	"testing"
)

func TestArgs(t *testing.T) {
	Register(Argument{
		name:         "arg",
		short:        "a",
		description:  "Test argument",
		defaultValue: "Test",
		values:       []string{"option1", "option2"},
		expectsValue: true,
	})
	Register(Argument{
		name:         "test",
		short:        "t",
		description:  "Test argument 2",
		expectsValue: false,
	})
	Register(Argument{
		name:         "no-short",
		short:        "",
		description:  "Example argument",
		expectsValue: true,
	})
	Register(Argument{
		name:        "examples-can-be-longer-and-longer",
		short:       "e",
		description: "Example argument 2",
	})

	fmt.Println("Registered Argument \"arg\"")
	Args["arg"] = "5"
	if Using("arg") {
		fmt.Println("Using Argument \"arg\".")
	}
	if Value("arg") != "" {
		fmt.Printf("\"arg\" has a value of \"%s\"\n", Value("arg"))
	} else {
		fmt.Println("\"arg\" has no value")
	}
	PrintUsage()
}
