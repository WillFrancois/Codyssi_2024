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

	for i := range len(lines) - 1 {
		value := strings.Split(lines[i], " ")
		num, err := strconv.Atoi(value[1])

		if err != nil {
			panic(err)
		}

		out += num
	}

	fmt.Println(out)
}
