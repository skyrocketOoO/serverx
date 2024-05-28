package main

import (
	"go-server-template/cmd"
)

func main() {
	cmd.Execute()
	myMap := map[string]string{
		"first key":  "first value",
		"second key": "second value",
		"third key":  "third value",
		"fourth key": "fourth value",
		"fifth key":  "fifth value",
	}
}
