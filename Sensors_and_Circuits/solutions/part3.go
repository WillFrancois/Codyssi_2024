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
	var results []string
	out := 0

	for len(lines) > 0 {

		fmt.Println(out)

		for i := range len(lines) {
			if lines[i] == "TRUE" {
				out += 1
			}
		}

		fmt.Println(lines)

		for i := 0; i < len(lines)-1; i += 2 {
			if i%4 == 0 {
				if lines[i] == "TRUE" && lines[i+1] == "TRUE" {
					results = append(results, "TRUE")
				} else {
					results = append(results, "FALSE")
				}
			} else {
				if lines[i] == "TRUE" || lines[i+1] == "TRUE" {
					results = append(results, "TRUE")
				} else {
					results = append(results, "FALSE")
				}
			}
			fmt.Println(results, i)
		}

		lines = results
		results = make([]string, 0)
		fmt.Println(lines)

	}

	fmt.Println(out)
}
