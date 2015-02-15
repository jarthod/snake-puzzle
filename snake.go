package main

import (
	"fmt"
)

// Definition of the puzzle, each piece in the chain is either angled or straight
const Size = 64
var Angled = [Size]int8{1, 0, 1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 1, 1}

type Vertex struct {
	X, Y, Z int8
}

type Solution struct {
	Pos [Size]Vertex
	Placed int
}

func printPuzzle(sol Solution) {
	var grid [4][4][4]int
	for i := 0; i < sol.Placed; i++ {
		pos := sol.Pos[i]
		grid[pos.X][pos.Y][pos.Z] = i+1
	}
	fmt.Print("        ")
	for z := 0; z < 4; z++ {
		fmt.Print("    z: ", z, "\t")
	}
	fmt.Print("\n")
	for y := 3; y >= 0; y-- {
		fmt.Print(" y: ", y, "   ")
		for z := 0; z < 4; z++ {
			for x := 0; x < 4; x++ {
				if grid[x][y][z] > 0 {
					fmt.Printf("%2d ", grid[x][y][z])
				} else {
					fmt.Print(" - ")
				}
			}
			fmt.Print("\t")
		}
		fmt.Print("\n")
	}
}

func main() {
	var sol Solution
	sol.Pos[0] = Vertex{0, 0, 0}
	sol.Pos[1] = Vertex{1, 0, 0}
	sol.Pos[2] = Vertex{2, 0, 0}
	sol.Pos[3] = Vertex{2, 1, 0}
	sol.Placed = 4
	printPuzzle(sol)
}