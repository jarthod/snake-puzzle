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
	Grid [4][4][4]int
	Placed int
}

func (v2 *Vertex) minus(v1 Vertex) Vertex {
	return Vertex{v2.X - v1.X, v2.Y - v1.Y, v2.Z - v1.Z}
}

func (v1 *Vertex) plus(v2 Vertex) Vertex {
	return Vertex{v2.X + v1.X, v2.Y + v1.Y, v2.Z + v1.Z}
}

func (sol *Solution) place(pos Vertex) bool {
	// fmt.Println("place:", sol.Placed+1, "@", pos)
	if pos.X < 0 || pos.X > 3 || pos.Y < 0 || pos.Y > 3 || pos.Z < 0 || pos.Z > 3 { return false; }
	if sol.Grid[pos.X][pos.Y][pos.Z] != 0 { return false; }
	sol.Pos[sol.Placed] = pos
	sol.Grid[pos.X][pos.Y][pos.Z] = sol.Placed + 1
	sol.Placed++
	return true
}

func (sol *Solution) removeLast() bool {
	if sol.Placed <= 0 { return false }
	pos := sol.Pos[sol.Placed-1]
	sol.Grid[pos.X][pos.Y][pos.Z] = 0
	sol.Placed--
	return true
}

func printPuzzle(sol Solution) {
	fmt.Print("        ")
	for z := 0; z < 4; z++ {
		fmt.Print("    z: ", z, "\t")
	}
	fmt.Print("\n")
	for y := 3; y >= 0; y-- {
		fmt.Print(" y: ", y, "   ")
		for z := 0; z < 4; z++ {
			for x := 0; x < 4; x++ {
				if sol.Grid[x][y][z] > 0 {
					fmt.Printf("%2d ", sol.Grid[x][y][z])
				} else {
					fmt.Print(" - ")
				}
			}
			fmt.Print("\t")
		}
		fmt.Print("\n")
	}
}

func solvePuzzle(sol *Solution) (bool) {
	if sol.Placed == Size {
		return true
	} else {
		last_pos := sol.Pos[sol.Placed-1]
		direction := last_pos.minus(sol.Pos[sol.Placed-2])
		// fmt.Println("last_pos:", last_pos, "direction:", direction, "angled:", Angled[sol.Placed-1])
		if Angled[sol.Placed-1] == 1 {
			choices := []Vertex{}
			if direction.X != 0 {
				choices = []Vertex{Vertex{0, 1, 0}, Vertex{0, -1, 0}, Vertex{0, 0, 1}, Vertex{0, 0, -1}}
			} else if direction.Y != 0 {
				choices = []Vertex{Vertex{1, 0, 0}, Vertex{-1, 0, 0}, Vertex{0, 0, 1}, Vertex{0, 0, -1}}
			} else {
				choices = []Vertex{Vertex{0, 1, 0}, Vertex{0, -1, 0}, Vertex{1, 0, 0}, Vertex{-1, 0, 0}}
			}
			for _, choice := range choices {
				if sol.place(last_pos.plus(choice)) {
					if solvePuzzle(sol) {
						return true
					} else {
						sol.removeLast();
					}
				}
			}
			return false
		} else {
			if sol.place(last_pos.plus(direction)) {
				if solvePuzzle(sol) {
					return true
				} else {
					sol.removeLast();
					return false
				}
			} else {
				return false
			}
		}
	}
}

func main() {
	var sol Solution
	sol.place(Vertex{0, 0, 0})
	sol.place(Vertex{1, 0, 0})
	sol.place(Vertex{2, 0, 0})
	sol.place(Vertex{2, 1, 0})
	fmt.Println("\n Start:")
	printPuzzle(sol)
	res := solvePuzzle(&sol);
	fmt.Println("\n Solved:", res)
	printPuzzle(sol)
}