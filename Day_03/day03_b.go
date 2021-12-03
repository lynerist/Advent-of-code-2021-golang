package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var oxygenRatingValues []string
	for sc.Scan(){
		oxygenRatingValues = append(oxygenRatingValues, sc.Text())
	}
	co2RatingValues := append([]string{}, oxygenRatingValues...) //A copy of oxygenRatingValues
	
	gamma, _ := strconv.ParseInt(filterValues(oxygenRatingValues, 0, true), 2, 64) //ParseInt converts from base 2 to decimal
	epsilon, _ := strconv.ParseInt(filterValues(co2RatingValues, 0, false), 2, 64)	
	fmt.Println(gamma*epsilon)
}

func filterValues(values []string, bitToConsider int, mostCommon bool)string{
	if len(values) == 1{ //Base case of the recursion, there is just the valid string
		return values[0]
	}
	//I divide all the remaining values in two slices looking to the considerated bit
	var valueWithZero, valueWithOne []string

	for _, value := range values{
		if rune(value[bitToConsider]) == '0'{
			valueWithZero = append(valueWithZero, value) 
		}else{
			valueWithOne = append(valueWithOne, value)
		}
	}
	//I apply the recursion keeping the slice with the most common bit if mostcommon is true
	if len(valueWithOne) >= len(valueWithZero) == mostCommon{
		return filterValues(valueWithOne, bitToConsider+1, mostCommon) 
	}
	return filterValues(valueWithZero, bitToConsider+1, mostCommon)    
}