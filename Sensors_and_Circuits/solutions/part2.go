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

	for i := 0; i < len(lines); i += 2 {
		if i%4 == 0 {
			if lines[i] == "TRUE" && lines[i+1] == "TRUE" {
				out += 1
			}
		} else {
			if lines[i] == "TRUE" || lines[i+1] == "TRUE" {
				out += 1
			}
		}
	}

	fmt.Println(out)
}
