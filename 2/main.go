package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

// Primary algorithm
func isSafe(report []int) bool {
	var forwardStepping = report[0] < report[1]

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if forwardStepping {
			if diff <= 0 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 || diff < -3 {
				return false
			}
		}
	}
	return true
}

// Data loading and processing
func stringsToInts(in []string) (out []int, err error) {
	var n int
	for _, i := range in {
		n, err = strconv.Atoi(i)
		if err != nil {
			return
		}
		out = append(out, n)
	}
	return
}

func loadReports() ([][]int, error) {
	var res [][]int

	reports := strings.Split(data, "\n")[0:1000]
	for _, r := range reports {
		r = strings.TrimSuffix(r, "\r")
		reportStrList := strings.Split(r, " ")

		reportIntList, err := stringsToInts(reportStrList)
		if err != nil {
			return nil, err
		}
		res = append(res, reportIntList)
	}

	return res, nil
}

func main() {

	reports, err := loadReports()
	if err != nil {
		panic(err)
	}

	var safe int

	for _, r := range reports {
		if isSafe(r) {
			safe++
		}
	}

	fmt.Printf("safe: %v\n", safe)
}
