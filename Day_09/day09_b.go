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

	var matrix [][]rune

	//I make a matrix of heights with a tall border
	for sc.Scan(){
		line := []rune{'9'}
		for _, point := range sc.Text(){
			line = append(line, point)
		}
		line = append(line, '9')
		matrix = append(matrix, line)
	}
	nines := make([]rune, len(matrix[0]))
	for i := range nines{
		nines[i] = '9'
	}
	matrix = append([][]rune{nines}, matrix... )
	matrix = append(matrix, nines)

	lowPoints := make(map[string]bool) //where we collect all low points

	for i, line := range matrix{
		for j, point := range line{
			if 	i!=0 && j!=0 && 
				i!=len(matrix)-1 && j!=len(line)-1 &&
				point < matrix[i+1][j] &&
				point < matrix[i-1][j] &&
				point < matrix[i][j+1] &&
				point < matrix[i][j-1]{
				lowPoints[fmt.Sprintf("%d,%d", i, j)] = true //This is a set of lowPoints
			}
		}
	} 

	var max, max2, max3 int //three max variables

	for point := range lowPoints{
		matrixCopy := append([][]rune{}, matrix...) //I need to preserve the original matrix 
		var x, y, bacinPoints int
		fmt.Sscanf(point, "%d,%d",&x,&y)
		_, bacinPoints = explore(matrixCopy, bacinPoints, x, y) //recursive exploration of the points

		if bacinPoints > max{
			max3 = max2
			max2 = max
			max = bacinPoints
		}else if bacinPoints > max2{
			max3 = max2
			max2 = bacinPoints
		}else if bacinPoints > max3{
			max3 = bacinPoints
		}
	}
	fmt.Println(max * max2 * max3)
}

func explore (matrix [][]rune, count, x, y int)([][]rune, int){
	if 	matrix[x][y] != '9'{ 	//If a point is equal to nine then it's a base case of the recursion
		count++					//I have to count all points different from nine
		matrix[x][y] = '9'		//I erase counted points making them nine
		matrix, count = explore(matrix,count, x+1,y)	//I recursively explore all near points
		matrix, count = explore(matrix,count, x-1,y)
		matrix, count = explore(matrix,count, x,y+1)
		matrix, count = explore(matrix,count, x,y-1)
	}
	return matrix, count
}