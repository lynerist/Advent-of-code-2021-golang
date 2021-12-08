package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var sum int
	for sc.Scan(){
		signal := strings.Split(sc.Text(), "|")
		rosettaStone := decode(strings.Fields(signal[0]))

		decodedDisplay := ""
		for _, digit := range strings.Fields(signal[1]){
			decodedDisplay += rosettaStone[sortSegments(digit)]
		}
		value, _ := strconv.Atoi(decodedDisplay)
		sum += value
	}
	fmt.Println(sum)
}

func decode(signal []string)map[string]string{
	rosettaStone := make(map[string]string)
	countIstances := make(map[rune]int)

	for _, digit := range signal{
		for _, segment := range digit{
			countIstances[segment]++
		}
	}

	for _, digit := range signal{
		switch len(digit){ 
		case 2: //There is just one digit with two segments
			rosettaStone[sortSegments(digit)] = "1"
		case 3: //There is just one digit with three segments
			rosettaStone[sortSegments(digit)] = "7"
		case 4: //There is just one digit with four segments
			rosettaStone[sortSegments(digit)] = "4"
		case 5:  //There are three digits with five segments
			for _, segment := range digit{
				if countIstances[segment] == 4 { //Only the digit 2 has a segment that has 4 instances in all digits (segment e)
					rosettaStone[sortSegments(digit)] = "2"
				}else if countIstances[segment] == 6{ //Only the digit 5 has a segment that has 6 instances in all digits (segment b)
					rosettaStone[sortSegments(digit)] = "5"
				}
			}
			if _, twoOrFive:= rosettaStone[sortSegments(digit)]; !twoOrFive{
				rosettaStone[sortSegments(digit)] = "3" //If it has 5 segments and it's not 2 or 5 then just digit 3 is left
			}
		case 6://There are three digits with six segments
			countSegmentsWithKeyInstances := make(map[int]int)
			for _, segment := range digit{
				countSegmentsWithKeyInstances[countIstances[segment]]++
			}
			if countSegmentsWithKeyInstances[4] == 0{      //There is just one digit with 6 segments missing the segment e 
				rosettaStone[sortSegments(digit)] = "9"    //(that has 4 instances in digits)
			}else if countSegmentsWithKeyInstances[7] == 1{//There is just one digit with 6 segments missing the segment d that is one of
				rosettaStone[sortSegments(digit)] = "0"    //the two segments with 7 instances in the ten digits
			}else{
				rosettaStone[sortSegments(digit)] = "6"
			}
		case 7://There is just one digit with seven segments
			rosettaStone[sortSegments(digit)] = "8"
		}
	}
	return rosettaStone
}

func sortSegments(digit string)(sortedDigit string){
	for _, segment := range []string{"a", "b", "c", "d", "e", "f", "g"}{
		if strings.Contains(digit, segment){
			sortedDigit += segment
		}
	}
	return
}
/*
digits with that number of segments
number of segment -> digits
len 2 -> 1
len 3 -> 7
len 4 -> 4
len 5 -> 2 3 5 
len 6 -> 0 6 9
len 7 -> 8

instances for segment: 
segment -> is present in n digits
a -> 8
b -> 6
c -> 8
d -> 7
e -> 4
f -> 9
g -> 7
*/