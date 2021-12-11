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

	var octopuses [][]int

	for sc.Scan(){
		line := []int{0}
		for _, octopus := range sc.Text(){
			line = append(line, int(octopus)-48) //fast rune-int conversion
		}
		line = append(line, 0)
		octopuses = append(octopuses, line)
	}
	zeros := make([]int, len(octopuses[0])) //I add a border of zeros
	for i := range zeros{
		zeros[i] = 0
	}
	octopuses = append([][]int{zeros}, octopuses... )
	octopuses = append(octopuses, zeros)
	
	var flashes int

	for step := 0; step<100; step++{
		//First I increase all octopuses counters	
		for y := 1; y<len(octopuses)-1; y++{
			for x := 1; x<len(octopuses[0])-1; x++{
				octopuses[y][x]++
			}
		}
		//Then I recursively let octopuses flash and charge their neightbours
		for y := range octopuses{
			for x := range octopuses[y]{
				flashes, octopuses = flash(y, x, flashes, octopuses)
			}
		}
	}
	fmt.Println(flashes)
}

func flash(y, x, flashes int, matrix [][]int)(int, [][]int){
	if matrix[y][x] > 9 {
		flashes++ // count flashes
		matrix[y][x] = 0
		for i:=-1; i<=1; i++ { //visit near octopuses
			for j:=-1; j<=1; j++ {
				if matrix[y+i][x+j] != 0 { //prevent double flash in a step
					matrix[y+i][x+j]++
				}
				flashes, matrix = flash(y+i, x+j, flashes, matrix)
			}
		}
	}
	return flashes, matrix
}