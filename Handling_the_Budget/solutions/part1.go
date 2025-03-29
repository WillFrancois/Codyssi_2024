package budget

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	var fileCon string = string(input)
	var numbers []string = strings.Split(fileCon, "\n")
	var out int64 = 0

	for i := range len(numbers) {
		if numbers[i] != "" {
			var value, err = strconv.Atoi(numbers[i])
			if err != nil {
				panic(err)
			} else {
				out += int64(value)
			}
		}
	}

	fmt.Println(out)
}
