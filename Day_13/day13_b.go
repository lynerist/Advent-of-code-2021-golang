package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	dots := make(map[string]bool) //A set of dots
	for sc.Scan(){
		if sc.Text() == ""{
			break
		}
		dots[sc.Text()] = true
	}

	for sc.Scan(){ //for each fold
		var foldingDirection rune
		var foldPoint int
		fmt.Sscanf(sc.Text(), "fold along %c=%d",&foldingDirection, &foldPoint)
		
		for dot := range dots{
			var x,y int
			fmt.Sscanf(dot, "%d,%d", &x,&y)
			if foldingDirection == 'y' && y >= foldPoint {
				y = foldPoint - (y-foldPoint)
				dots[fmt.Sprintf("%d,%d", x,y)] = true //both new or overlapped dots
			}else if foldingDirection == 'x' && x >= foldPoint {
				x = foldPoint - (x-foldPoint)
				dots[fmt.Sprintf("%d,%d", x,y)] = true //both new or overlapped dots
			}
		}
	}

	for y:= 0; y<6; y++{ //I print the code
		for x:=0; x<8*4+7; x++ { //8 letters * width 4 + 7 spaces
			if dots[fmt.Sprintf("%d,%d", x,y)]{
				fmt.Print("#")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}