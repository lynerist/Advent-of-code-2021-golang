package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	var lanternfishesByWaitingTime [9]int
	//I collect the number of fishes by age
	for _, fish := range strings.Split(sc.Text(), ","){
		fishWaitingTime, _ := strconv.Atoi(fish)
		lanternfishesByWaitingTime[fishWaitingTime]++
	}
	//Here there is the grown process
	for generation := 0; generation < 80; generation++{
		justBred := lanternfishesByWaitingTime[0]
		for daysToWait := range lanternfishesByWaitingTime[:len(lanternfishesByWaitingTime)-1]{
			lanternfishesByWaitingTime[daysToWait] = lanternfishesByWaitingTime[daysToWait+1]  
		}
		lanternfishesByWaitingTime[6] += justBred //start waiting other 6 days
		lanternfishesByWaitingTime[8] = justBred //new fishes
	}  

	var numberOfFishes int
	for _, fishes := range lanternfishesByWaitingTime{
		numberOfFishes += fishes
	}
	fmt.Println(numberOfFishes)
}