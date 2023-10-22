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
		Name:         "arg",
		Short:        "a",
		Description:  "Test argument",
		DefaultValue: "Test",
		Values:       []string{"option1", "option2"},
		ExpectsValue: true,
	})
	Register(Argument{
		Name:         "test",
		Short:        "t",
		Description:  "Test argument 2",
		ExpectsValue: false,
	})
	Register(Argument{
		Name:         "no-Short",
		Short:        "",
		Description:  "Example argument",
		ExpectsValue: true,
	})
	Register(Argument{
		Name:        "examples-can-be-longer-and-longer",
		Short:       "e",
		Description: "Example argument 2",
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
