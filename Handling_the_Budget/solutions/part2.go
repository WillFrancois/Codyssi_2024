package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	var fileCon string = string(input)
	var numbers []string = strings.Split(fileCon, "\n")
	var out int64 = 0

	sort.Slice(numbers, func(i, j int) bool {
		num1, _ := strconv.Atoi(numbers[j])
		num2, _ := strconv.Atoi(numbers[i])
		return num1 < num2
	})

	fmt.Println(numbers)

	for i := 20; i < len(numbers); i++ {
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
