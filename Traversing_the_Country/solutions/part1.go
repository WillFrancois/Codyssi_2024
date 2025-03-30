package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")
	location_mapping := make(map[string]bool, 0)

	for i := range len(lines) - 1 {
		locations := strings.Split(string(lines[i]), " <-> ")
		location_mapping[string(locations[0])] = true
		location_mapping[string(locations[1])] = true
	}

	fmt.Println(len(location_mapping))
}
