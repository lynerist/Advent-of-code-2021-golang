package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var numberOfIncreases int
	
	//Read first element
	sc.Scan()
	prevDepth, _ := strconv.Atoi(sc.Text())

	//For each other element we comprare it with the previous one
	for sc.Scan(){
		depth, _ := strconv.Atoi(sc.Text())
		if depth > prevDepth{
			numberOfIncreases++
		}
		prevDepth = depth
	}

	fmt.Println(numberOfIncreases)
}