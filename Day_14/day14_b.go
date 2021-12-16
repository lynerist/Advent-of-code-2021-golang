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

	sc.Scan()
	polymer := sc.Text()
	sc.Scan()
	insertionRules := make(map[string]string)

	for sc.Scan(){
		var input, insertion string
		fmt.Sscanf(sc.Text(), "%s -> %s", &input, &insertion)
		insertionRules[input] = insertion
	}

	pairs := make(map[string]int)
	countElements := map[string]int{polymer[0:1]:1}
	for i:=0; i<len(polymer)-1; i++{
		pairs[polymer[i:i+2]]++
		countElements[polymer[i+1:i+2]]++
	}
	
	for i:=0; i<40; i++{
		newPairs := make(map[string]int)
		for pair, instances := range pairs{
			newPairs[pair[0:1] + insertionRules[pair]] += instances
			newPairs[insertionRules[pair] + pair[1:2]] += instances
			countElements[insertionRules[pair]] += instances
		}
		pairs = newPairs
	}

	min, max := countElements[polymer[0:1]], countElements[polymer[0:1]]

	for _, count := range countElements{
		if count > max{
			max = count
		} else if count < min {
			min = count
		}
	}
	fmt.Println(max-min)
}
