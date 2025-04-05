package main

import (
	"fmt"
	"os"
	"strings"
)

func countUniqueElements(input_array []string) int {
	counter := make(map[string]int, len(input_array))
	out := 0

	for i := range len(input_array) {
		counter[input_array[i]] += 1
	}

	for range len(counter) {
		out += 1
	}

	return out
}

func main() {
	contents, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")
	location_mapping := make(map[string][]string, 1)
	location_mapping["STT"] = append(location_mapping["STT"], "STT")
	updated_locations := make([]string, 0)

	for range 3 {
		for i := range len(lines) {
			locations := strings.Split(string(lines[i]), " <-> ")
			cur_len := len(location_mapping["STT"])

			if len(locations) < 2 {
				continue
			}

			for j := range cur_len {
				if locations[0] == location_mapping["STT"][j] && len(location_mapping[locations[1]]) == 0 {
					updated_locations = append(updated_locations, locations[1])
					location_mapping[locations[1]] = make([]string, 1)
					location_mapping[locations[1]][0] = "STT"
				} else if locations[1] == location_mapping["STT"][j] && len(location_mapping[locations[0]]) == 0 {
					updated_locations = append(updated_locations, locations[0])
					location_mapping[locations[0]] = make([]string, 1)
					location_mapping[locations[0]][0] = "STT"
				}
			}
		}
		location_mapping["STT"] = append(location_mapping["STT"], updated_locations...)
	}

	fmt.Println(countUniqueElements(location_mapping["STT"]))
}
