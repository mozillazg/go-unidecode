package unidecode_test

import (
	"fmt"

	"github.com/mozillazg/go-unidecode"
)

func ExampleUnidecode() {
	s := "北京"
	fmt.Println(unidecode.Unidecode(s))
	// Output: Bei Jing
}

func ExampleUnidecode_ascii() {
	s := "abc"
	fmt.Println(unidecode.Unidecode(s))
	// Output: abc
}
