package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

// Displacement is the x axis and depth is the y axis
// aim rapresents the inclination of the submarine
// a positive aim means that the submarine is pointing toward the seabed
	var depth, displacement, aim int
		
	for sc.Scan(){
		command := strings.Fields(sc.Text())
		delta, _ := strconv.Atoi(command[1])

		switch command[0]{
		case "forward":
			displacement += delta
			depth += aim * delta
		case "down":
			aim += delta
		case "up":
			aim -= delta
		}
	}
	fmt.Println(depth * displacement)
}