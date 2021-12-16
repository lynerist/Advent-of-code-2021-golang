package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct{
	risk int
	distance int
	predecessor string
}

func strip(y, x int)string{
	return fmt.Sprintf("%d-%d", y, x)
}

func unStrip(position string)(y, x int){
	fmt.Sscanf(position, "%d-%d", &y, &x)
	return 
}

func getNextNode(unvisited map[string]bool, state map[string]Node)(nextNode string){
	minDistance := 0
	for node := range unvisited{
		if state[node].distance < minDistance || nextNode == ""{
			nextNode = node
			minDistance = state[node].distance
		}
	}
	return
}

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	unvisited := make(map[string]bool)
	visited := make(map[string]bool)
	state := make(map[string]Node)

	var numberOfLines, numberOfColumns int
	for sc.Scan(){
		numberOfColumns = 0
		for _, risk := range sc.Text(){
			node := strip(numberOfLines, numberOfColumns)
			unvisited[node] = true
			state[node] = Node{int(risk)-48, math.MaxInt64, ""}
			numberOfColumns++
		}
		numberOfLines++
	}
	state["0-0"] = Node{state["0-0"].risk, 0, ""}

	endNode := strip(numberOfLines-1, numberOfColumns-1)
	for len(unvisited) != 0 && state[endNode].distance == math.MaxInt64{
		currentNode := getNextNode(unvisited, state)
		y, x := unStrip(currentNode)

		var neighbours []string
		if y > 0 {
			neighbours = append(neighbours, strip(y-1, x))
		}
		if y < numberOfLines-1{
			neighbours = append(neighbours, strip(y+1, x))
		}
		if x > 0 {
			neighbours = append(neighbours, strip(y, x-1))
		}
		if x < numberOfColumns-1{
			neighbours = append(neighbours, strip(y, x+1))
		}

		for _, neighbour := range neighbours {
			if state[currentNode].distance + state[neighbour].risk < state[neighbour].distance {
				newState := state[neighbour]
				newState.distance = state[currentNode].distance + state[neighbour].risk
				newState.predecessor = currentNode
				state[neighbour] = newState
			}
		}
		visited[currentNode] = true
		delete(unvisited, currentNode)
	}
	fmt.Println(state[endNode].distance)
}

