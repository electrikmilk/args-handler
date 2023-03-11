/*
 * Copyright (c) 2023 Brandon Jordan
 */

package args

import (
	"fmt"
	"testing"
)

func TestArgs(t *testing.T) {
	Register("arg", "a", "Test argument", true)
	Register("test", "t", "Test argument 2", false)
	Register("example", "e", "Example argument", true)
	fmt.Println("Registered argument \"arg\"")
	if Using("arg") {
		fmt.Println("Using argument \"arg\".")
	}
	if Value("arg") != "" {
		fmt.Printf("\"arg\" has a value of \"%s\"\n", Value("arg"))
	} else {
		fmt.Println("\"arg\" has no value")
	}
}
