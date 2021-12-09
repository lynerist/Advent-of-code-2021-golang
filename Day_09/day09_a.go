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

	//I make a matrix of heights with a tall border
	var matrix [][]rune
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

	riskLevel := 0

	for i, line := range matrix{
		for j, point := range line{
			if 	i!=0 && j!=0 && 
				i!=len(matrix)-1 && j!=len(line)-1 &&
				point < matrix[i+1][j] &&
				point < matrix[i-1][j] &&
				point < matrix[i][j+1] &&
				point < matrix[i][j-1]{
				riskLevel += int(point)-48 +1 //digit-48 is the int conversion of the digit (ASCII code of 0 is 48)
			}
		}
	} 
	fmt.Println(riskLevel)
}