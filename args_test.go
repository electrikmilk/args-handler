/*
 * Copyright (c) 2022 Brandon Jordan
 */

package args

import (
	"fmt"
	"testing"
)

func TestArgs(t *testing.T) {
	registerArg("arg", "a", "Test argument")
	fmt.Println("Registered argument \"arg\"")
	if arg("arg") {
		fmt.Println("Using argument \"arg\".")
	}
	if argValue("arg") != "" {
		fmt.Printf("\"arg\" has a value of \"%s\"\n", argValue("arg"))
	} else {
		fmt.Println("\"arg\" has no value")
	}
	usage()
}
