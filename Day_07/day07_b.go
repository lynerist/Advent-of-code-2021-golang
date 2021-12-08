package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	values := strings.Split(sc.Text(), ",")
	
	mean := 0.0
	for _, position := range values{
		positionValue, _ := strconv.Atoi(position)
		mean += float64(positionValue)
	}
	mean /= float64(len(values))
	
	allignPoint := int(mean)

	fuelCost := 0
	for _, position := range values{
		positionValue, _ := strconv.Atoi(position)
		fuelCost+=fuel(int(math.Abs(float64(positionValue -allignPoint))))
	}
	fmt.Println(fuelCost)
}

func fuel(n int)int{
	if n == 0{
		return 0
	}
	return n + fuel(n-1)
}