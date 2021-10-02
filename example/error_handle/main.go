package main

import (
	"github.com/cncamp/example/error_handle/parse"
)

func main() {
	var examples = []string{
		"1, 2, 3, 4, 5",
		"100, 2.45, 20, 1.22",
		"hello world",
	}
	parse.Parse(examples)
}
