package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	errorPoints := map[rune]int{')':1, ']':2, '}':3, '>':4}
	closingCharacters := map[rune]rune{'(':')', '[':']', '{':'}', '<':'>'}
	var scores []int

	LINE: for sc.Scan(){
		var score int
		var stack []rune
		for _, char := range sc.Text(){
			switch char{
			case '(', '[', '{', '<':
				stack = append([]rune{closingCharacters[char]}, stack...) //Append elements to the top of the stack
			case ')', ']', '}', '>':
				if len(stack) == 0 || stack[0] != char{
					continue LINE //We ignore corrupted lines
				}else{
					stack = stack[1:] //Remove the top element of the stack
				}
			}
		}
		for _, char := range stack{
			score = score * 5 + errorPoints[char] 
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
