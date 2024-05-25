package main

import "fmt"

func intToRoman(num int) string {
	thousands_map := map[int]string {
		1000: "M",
		2000: "MM",
		3000: "MMM",
	}

	hundreds_map := map[int]string {
		100: "C",
		200: "CC",
		300: "CCC",
		400: "CD",
		500: "D",
		600: "DC",
		700: "DCC",
		800: "DCCC",
		900: "CM",
	}

	tens_map := map[int]string {
		10: "X",
		20: "XX",
		30: "XXX",
		40: "XL",
		50: "L",
		60: "LX",
		70: "LXX",
		80: "LXXX",
		90: "XC",
	}

	units_map := map[int]string {
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}

    thousands := num - num%1000
	num = num - thousands
	hundreds := num  - num%100
	num = num -hundreds
	tens := num - num%10
	num = num - tens
	units := num

	s := fmt.Sprintf("%s%s%s%s", thousands_map[thousands], hundreds_map[hundreds], tens_map[tens], units_map[units])

	return s
	
}


func main() {
	num := 1994

	answer := intToRoman(num)

	fmt.Println(answer)

}