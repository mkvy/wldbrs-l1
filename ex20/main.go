package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(in string) string {
	words := strings.Split(in, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	fmt.Println("--- After reverse: ---")
	fmt.Println(reverseString(str))
}
