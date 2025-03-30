package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")
	out := 0

	mapping := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#"
	fmt.Println(len(mapping))

	for i := range len(lines) - 1 {
		value := strings.Split(lines[i], " ")
		base, err := strconv.Atoi(value[1])

		if err != nil {
			panic(err)
		}

		increment, err := strconv.ParseInt(value[0], base, 0)

		if err != nil {
			panic(err)
		}

		out += int(increment)
	}

	result := ""

	for out > 0 {
		result = string(mapping[out%65]) + result
		out /= 65
	}

	fmt.Println(result)
}
