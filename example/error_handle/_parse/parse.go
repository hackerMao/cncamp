package parse

import (
	"fmt"
	"strconv"
	"strings"
)

/**
自定义包实现者应该遵守的最佳实践:
1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
2）向包的调用者返回错误值（而不是 panic）。
*/

type Error struct {
	Index int
	Word  string
	Err   error
}

func (e *Error) String() string {
	return fmt.Sprintf("pkg _parse: error parsing %q as int", e.Word)
}

func Parse(input string) (numbers []int, err error) {
	defer func() {
		if err := recover(); err != nil {
			var ok bool
			err, ok = err.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", err)
			}
		}
	}()
	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("the length of fields is 0")
	}
	for index, field := range fields {
		number, err := strconv.Atoi(field)
		if err != nil {
			panic(&Error{index, field, err})
		}
		numbers = append(numbers, number)
	}
	return
}
