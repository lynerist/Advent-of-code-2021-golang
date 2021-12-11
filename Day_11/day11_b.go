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
			line = append(line, int(octopus)-48)
		}
		line = append(line, 0)
		octopuses = append(octopuses, line)
	}
	zeros := make([]int, len(octopuses[0]))
	for i := range zeros{
		zeros[i] = 0
	}
	octopuses = append([][]int{zeros}, octopuses... )
	octopuses = append(octopuses, zeros) //I add a border of zeros
	
	var steps int
	var allFlashed bool

	for steps = 0; !allFlashed; steps++{	
		//First I increase all octopuses counters
		for y := 1; y<len(octopuses)-1; y++{
			for x := 1; x<len(octopuses[0])-1; x++{
				octopuses[y][x]++
			}
		}
		//Then I recursively let octopuses flash and charge their neightbours
		for y := range octopuses{
			for x := range octopuses[y]{
				octopuses = flash(y, x, octopuses)
			}
		}
		//Finally I see if all of them has flashed
		allFlashed = true
		for y := 1; allFlashed && y<len(octopuses)-1; y++{
			for x := 1; allFlashed && x<len(octopuses[0])-1; x++{
				allFlashed = octopuses[y][x] == 0
			}
		}
	}
	fmt.Println(steps)
}

func flash(y, x int, matrix [][]int)[][]int{
	if matrix[y][x] > 9 {
		matrix[y][x] = 0
		for i:= -1; i<=1; i++ { //visit near octopuses
			for j:=-1; j<=1; j++ {
				if matrix[y+i][x+j] != 0{ //prevent double flash in a step
					matrix[y+i][x+j]++
				}
				matrix = flash(y+i, x+j, matrix)
			}
		}
	}
	return matrix
}