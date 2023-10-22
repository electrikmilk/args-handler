/*
 * Copyright (c) 2023 Brandon Jordan
 */

package args

import (
	"fmt"
	"os"
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
	fmt.Println("Registered Argument \"arg\"")
	Register(Argument{
		Name:         "test",
		Short:        "t",
		Description:  "Test argument 2",
		ExpectsValue: false,
	})
	fmt.Println("Registered Argument \"test\"")
	Register(Argument{
		Name:         "no-short",
		Short:        "",
		Description:  "Example argument",
		ExpectsValue: true,
	})
	fmt.Println("Registered Argument \"no-short\"")
	Register(Argument{
		Name:        "examples-can-be-longer-and-longer",
		Short:       "e",
		Description: "Example argument 2",
	})
	fmt.Print("Registered Argument \"long-example\"\n\n")

	os.Args[1] = "--arg=test"
	os.Args[2] = "-e"

	parseArgs()

	for _, arg := range registered {
		if Using(arg.Name) {
			fmt.Printf("Using argument \"%s\"", arg.Name)
		} else {
			fmt.Printf("NOT using argument \"%s\"", arg.Name)
		}

		fmt.Print("\n")

		if Value(arg.Name) != "" {
			fmt.Printf("\"%s\" has a value of \"%s\"\n", arg.Name, Value("arg"))
		} else {
			fmt.Printf("\"%s\" has NO value\n", arg.Name)
		}

		fmt.Print("\n")
	}

	PrintUsage()
}
