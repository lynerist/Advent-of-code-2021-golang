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

	connections := make(map[string][]string)

	for sc.Scan(){
		path := strings.Split(sc.Text(), "-")
		connections[path[0]] = append(connections[path[0]], path[1])
		connections[path[1]] = append(connections[path[1]], path[0])
	}
	fmt.Println(explore("start", connections, map[string]bool{}, true))
}

func explore(start string, connections map[string][]string, alreadyVisited map[string]bool, exception bool)int{
	if start == "end"{
		return 1
	}
	visited := map[string]bool{start:strings.ToLower(start)==start}
	for node := range alreadyVisited{
		visited[node] = alreadyVisited[node]
	}
	var paths int
	for _, node := range connections[start]{
		if ! visited[node]{
			paths += explore(node, connections, visited, exception)
		}else if exception && node != "start"{
			paths += explore(node, connections, visited, false)
		}
	} 
	return paths
}