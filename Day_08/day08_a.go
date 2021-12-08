package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var tenUniquePatterns [][]string
	var fourDigitValues [][]string

	for sc.Scan(){
		signal := strings.Split(sc.Text(), "|")
		tenUniquePatterns = append(tenUniquePatterns, strings.Fields(signal[0]))
		fourDigitValues = append(fourDigitValues, strings.Fields(signal[1]))
	}

	var istances int

	for _, display := range fourDigitValues{
		for _,digit := range display{
			if len(digit)==2 || len(digit)==3 || len(digit)==4 || len(digit)==7{
				istances++
			}
		}
	}

	fmt.Println(istances)
}