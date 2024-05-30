package main

import (
	"go-server-template/cmd"
)

func main() {
	cmd.Execute()
}

func Foo(bar string, baz string) {}

// A newline before or after an element requires newlines for the opening and
// closing braces.
var ints = []int{
	1, 2,
	3, 4,
}

// A newline between consecutive elements requires a newline between all
// elements.
var matrix = [][]int{
	{1},
	{2},
	{
		3,
	},
}

var V interface{} = 3

type T struct{}

func F()
