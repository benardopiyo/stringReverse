package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Welcome"
	fmt.Println(reverse(str1))
	fmt.Println(len(str1))
	fmt.Println(len(reverse(str1)))
}

func reverse(str string) string {
	var output strings.Builder
	for _, word := range strings.Split(str, " ") {
		if len(word) >= 5 {
			ru := []rune(word)
			for i, j := 0, len(ru)-1; i < j; i, j = i+1, j-1 {
				ru[i], ru[j] = ru[j], ru[i]
			}
			output.WriteString(string(ru))
			output.WriteString(" ")
		} else {
			output.WriteString(word)
			output.WriteString(" ")
		}
	}
	return strings.TrimRight(output.String(), " ")
}
