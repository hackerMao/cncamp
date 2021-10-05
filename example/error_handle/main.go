package main

import (
	"fmt"
	parse "github.com/cncamp/example/error_handle/_parse"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 2.45 20 1.22",
		"hello world",
	}
	for _, example := range examples {
		numbers, err := parse.Parse(example)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(numbers)
	}
}
