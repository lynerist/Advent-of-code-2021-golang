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
	//We store how many lines pass through a point
	linesInPoint := make(map[string]int)

	for sc.Scan(){
		var start, end [2]int //points are rappresented as arrays of 2 int [x, y]
		fmt.Sscanf(sc.Text(), "%d,%d -> %d,%d", &(start[0]),&(start[1]), &(end[0]), &(end[1]))
		
		linesInPoint[fmt.Sprintf("%d,%d",start[0], start[1])]++ //start is added manually
		for start[0] != end[0] || start[1] != end[1]{
			//We go ahead by one point near to the end
			if start[0] > end[0]{
				start[0]--
			}else if start[0] < end[0]{
				start[0]++
			}
			if start[1] > end[1]{
				start[1]--
			}else if start[1] < end[1]{
				start[1]++
			}
			//We increase the number of lines passing through the current point by one
			linesInPoint[fmt.Sprintf("%d,%d",start[0], start[1])]++ 
		}
	}

	var countPoints int
	for _, v := range linesInPoint{
		if v>1{ //We count all points with two or more lines passing through them
			countPoints++
		}
	}
	fmt.Println(countPoints)
}