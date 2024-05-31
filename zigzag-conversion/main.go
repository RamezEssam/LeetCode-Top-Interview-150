package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := []strings.Builder{}
	for i:= 0; i < numRows; i++ {
		rows = append(rows, strings.Builder{})
	}

	height := 0
	travelDown := false
	for i:= 0; i < len(s); i++ {
		if height == 0 || height == numRows -1 {
			travelDown = !travelDown
		}

		rows[height].WriteByte(s[i])

		if travelDown {
			height += 1
		}else {
			height -= 1
		}

	}

	answer := strings.Builder{}

	for _, row := range rows {
		answer.WriteString(row.String())
	}

	return answer.String()
}

func main() {
	input := "PAYPALISHIRING"
	numRows := 4

	answer := convert(input, numRows)

	fmt.Println(answer)
}
