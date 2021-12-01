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
	
	//Read first three elements
	sc.Scan()
	prevPrevPrevDepth, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	prevPrevDepth, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	prevDepth, _ := strconv.Atoi(sc.Text())

	//For each element we compare it with the one that was three places before it
	//That's because each couple of triplet that we have to compare shares two elements
	//So we compare just the different one (in this example B > A if and only if z > w)
	// w A
	// x A B
	// y A B 
	// z   B 
	for sc.Scan(){
		depth, _ := strconv.Atoi(sc.Text())
		if depth > prevPrevPrevDepth{
			numberOfIncreases++
		}
		prevPrevPrevDepth = prevPrevDepth
		prevPrevDepth = prevDepth
		prevDepth = depth
	}

	fmt.Println(numberOfIncreases)
}