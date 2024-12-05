package main

// Solution uses regex, not exactly proud of it...

import (
	_ "embed"
	"fmt"
	"regexp"
)

//go:embed data.txt
var data string

func partone() int {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	items := r.FindAllString(data, -1)
	var sum int

	for _, i := range items {
		var a, b int
		if r, err := fmt.Sscanf(i, "mul(%d,%d)", &a, &b); r != 2 || err != nil {
			fmt.Printf("r == %d\nerr == %s\n", r, err)
			panic(err)
		}
		sum += a * b
	}

	return sum
}

func extractMulParams(s string) (int, int) {
	var a, b int
	if r, err := fmt.Sscanf(s, "mul(%d,%d)", &a, &b); r != 2 || err != nil {
		fmt.Printf("r == %d\nerr == %s\n", r, err)
		panic(err)
	}

	return a, b
}

func parttwo() int {
	r := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	var enabled = true
	items := r.FindAllString(data, -1)

	var sum int

	for _, i := range items {
		if i == "do()" {
			enabled = true
		} else if i == "don't()" {
			enabled = false
		} else {
			if enabled {
				a, b := extractMulParams(i)
				sum += a * b
			}
		}
	}
	return sum
}

func main() {
	fmt.Printf("partone: %v\n", partone())

	fmt.Printf("parttwo: %v\n", parttwo())
}
