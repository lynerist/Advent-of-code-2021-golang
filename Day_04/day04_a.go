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

	sc.Scan()
	toDraw := strings.Split(sc.Text(), ",")
	//I prepare a map to associate to each number the turn where it's drawn
	turnToDraw := make(map[string]int)
	for turn, number := range toDraw{
		turnToDraw[number] = turn
	}
	//Skip empty line
	sc.Scan()

	var bestBoard [][]string
	winningTurnOfBest := len(turnToDraw)

	var currentBoard[][]string
	for {
		sc.Scan()
		if sc.Text() == "" {
			if len(currentBoard)==0{
				break
			}

			winningTurn := len(turnToDraw)

			var linesLastDrowsTurns, columnsLastDrowsTurns [5]int
			for i:=0; i<len(currentBoard); i++{
				for j:=0; j<len(currentBoard); j++{
					cellTurn := turnToDraw[currentBoard[i][j]]
					if cellTurn > linesLastDrowsTurns[i]{
						linesLastDrowsTurns[i] = cellTurn
					}
					if cellTurn > columnsLastDrowsTurns[j]{
						columnsLastDrowsTurns[j] = cellTurn
					}
				}
			}
			//I find the winning turn as a search of minimum between the lastDrowTurns of columns and lines
			for _, turn := range linesLastDrowsTurns {
				if turn < winningTurn {
					winningTurn = turn
				}
			}
			for _, turn := range columnsLastDrowsTurns {
				if turn < winningTurn {
					winningTurn = turn
				}
			}
			//The best board has the minimum winning turn
			if winningTurn < winningTurnOfBest{
				bestBoard = currentBoard
				winningTurnOfBest = winningTurn
			}
			//I reset the board
			currentBoard = [][]string{}
		}else{
			currentBoard = append(currentBoard, strings.Fields(sc.Text()))
		}
	}

	//I calculate the score
	var sum int
	for _, line := range bestBoard{
		for _, number := range line{
			intNumber, _ := strconv.Atoi(number)
			//Unmarked numbers
			if turnToDraw[number] > winningTurnOfBest{
				sum += intNumber
			}
		}
	}

	winningNumber, _ := strconv.Atoi(toDraw[winningTurnOfBest])
	fmt.Println(sum * winningNumber)
}