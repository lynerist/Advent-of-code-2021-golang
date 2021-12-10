package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var score int
	errorPoints := map[rune]int{')':3, ']':57, '}':1197, '>':25137}
	closingCharacters := map[rune]rune{'(':')', '[':']', '{':'}', '<':'>'}
	
	LINE: for sc.Scan(){
		var stack []rune
		for _, char := range sc.Text(){
			switch char{
			case '(', '[', '{', '<':
				stack = append([]rune{closingCharacters[char]}, stack...) //Append elements to the top of the stack
			case ')', ']', '}', '>':
				if len(stack) == 0 || stack[0] != char{
					score += errorPoints[char]
					continue LINE //Stop after the first error. I don't use a simple break due to the fact that we are in a switch
				}else{
					stack = stack[1:] //Remove the top element of the stack
				}
			}
		}	
	}
	fmt.Println(score)
}

