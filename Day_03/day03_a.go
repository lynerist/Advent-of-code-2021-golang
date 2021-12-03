package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var sumByColumns [12]int
	var numberOfLines int

	for sc.Scan(){
		for i, bit := range sc.Text(){
			sumByColumns[i]+= int(bit-48) // '0' has ASCII code 48 so '0'- 48 = 0 and '1'- 48 = 1
		}
		numberOfLines++
	}

	var gamma, epsilon int
	for i, sum := range sumByColumns{ 
		if sum > numberOfLines/2{ // if sum is greater then half of the numbers of line it means that '1' is the most common bit
			gamma += int(math.Pow(2, float64(len(sumByColumns)-1-i))) //I cumulate the current power of two for converting from binary to decimal
		}else{
			epsilon += int(math.Pow(2, float64(len(sumByColumns)-1-i)))
		}
	}
	fmt.Println(gamma*epsilon)
}