package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file_bytes, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file_bytes), "\n")
	out := 0

	for i := range len(lines) {
		if lines[i] == "TRUE" {
			out += i + 1
		}
	}

	fmt.Println(out)
}
