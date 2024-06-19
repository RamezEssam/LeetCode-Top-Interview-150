package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
    output := []string{}

	rem := maxWidth
	line := strings.Builder{}
	line_begin := 0
	num_chars := 0 
	for i:=0; i < len(words)-1; i++ {

		rem -= (len(words[i])+1)
		num_chars += len(words[i])

		if (rem - len(words[i+1])) < 0 {
			num_space_chars := maxWidth - num_chars
			num_spaces := i - line_begin 
			space_chars := 0
			if num_spaces == 0 {
				line.WriteString(words[line_begin])
				for i:=0; i < num_space_chars; i++ {
					line.WriteRune(' ')
				}
				output = append(output, line.String())
				rem = maxWidth
				line = strings.Builder{}
				line_begin = i+1
				num_chars = 0
				continue 
			}else if num_space_chars % num_spaces == 0 {
				space_chars = num_space_chars / num_spaces
			}else{
				space_chars = (num_space_chars / num_spaces) +1
			}

			for j:= line_begin; j < i+1; j++ {
				line.WriteString(words[j])
				if num_spaces > 0 {
					for k:=0; k < space_chars; k++ {
						line.WriteRune(' ')
						num_space_chars --
					}
					num_spaces --
					if num_spaces > 0 {
						if num_space_chars % num_spaces == 0 {
							space_chars = num_space_chars / num_spaces
						}else{
							space_chars = (num_space_chars / num_spaces) +1
						}
					}
				}
				
			}
			output = append(output, line.String())
			rem = maxWidth
			line = strings.Builder{}
			line_begin = i+1
			num_chars = 0 
		}
	}

	last_line := strings.Builder{}
	num_chars_written := 0
	for i:=line_begin; i < len(words); i++ {
		if i != len(words)-1 {
			last_line.WriteString(words[i])
			last_line.WriteRune(' ')
			num_chars_written += (len(words[i]) +1)
		}else{
			last_line.WriteString(words[i])
			num_chars_written += len(words[i])
		}
	}

	num_spaces := maxWidth - num_chars_written
	for i:=0; i < num_spaces; i++ {
		last_line.WriteRune(' ')
	}

	output = append(output, last_line.String())

	return output
}

func main() {
	words := []string{"Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"}

	output := fullJustify(words, 20)

	for _,line := range output {
		fmt.Printf("%s, len = %d\n", line, len(line))
	}

}