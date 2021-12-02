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

	//I need just two variables, one for each coordinate
	var depth, displacement int
		
	for sc.Scan(){
		command := strings.Fields(sc.Text())
		delta, _ := strconv.Atoi(command[1])

		switch command[0]{
		case "forward":
			displacement += delta
		case "down":
			depth += delta
		case "up":
			depth -= delta
		}
	}
	fmt.Println(depth * displacement)
}