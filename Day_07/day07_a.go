package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	
	
	var positions []int
	for _, value := range strings.Split(sc.Text(), ","){
		position, _ := strconv.Atoi(value)
		positions = append(positions, position)
	}
	sort.Ints(positions)

	alignPoint := positions[len(positions)/2]

	fuelCost := 0
	for _, position := range positions{
		fuelCost += int(math.Abs(float64(position-alignPoint)))
	}
	fmt.Println(fuelCost)
}