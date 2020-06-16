package main

import (
	"fmt"
	"time"
)

type board [][]int

type Life struct {
	curr, next board
	size       int
}

const LiveCell = 1
const DeadCell = 0

func (gol *Life) printCurrentBoard() {
	fmt.Print("\n")
	for i := range gol.curr {
		for j := range gol.curr[i] {
			if gol.curr[i][j] == LiveCell {
				fmt.Print("^ ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (gol *Life) countLiveNbrs(row, col int) int {
	var directions = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}
	liveNbors := 0
	for _, v := range directions { // count the live neighbors
		nrow := v[0] + row
		ncol := col + v[1]

		if nrow >= 0 && nrow < gol.size && ncol >= 0 && ncol < gol.size {
			if gol.curr[nrow][ncol] == LiveCell {
				liveNbors++
			}
		}
	}
	return liveNbors
}

func (gol *Life) applyLifeRules(row, col int) {
	liveNbors := gol.countLiveNbrs(row, col)
	// apply the game of life rules
	if gol.curr[row][col] == LiveCell && (liveNbors < 2 || liveNbors > 3) {
		gol.next[row][col] = DeadCell
	}
	if gol.curr[row][col] == LiveCell && (liveNbors == 2 || liveNbors == 3) {
		gol.next[row][col] = LiveCell
	}
	if gol.curr[row][col] == DeadCell && liveNbors == 3 {
		gol.next[row][col] = LiveCell
	}
}

// evolves gol with everu tick by applying the rules
// after rules are applied to each cell, the current state is saved into a new board we call next
// once all cell are resolved, the current state is swapped with the next state
func (gol *Life) evolve(tick time.Duration) {
	if gol.size < 1 {
		fmt.Print("size of board cannot be less than 1")
		return
	}
	gol.printCurrentBoard()
	for range time.Tick(tick) {
		for i := range gol.curr {
			for j := range gol.curr[i] {
				gol.applyLifeRules(i, j)
			}
		}
		tmp := gol.curr
		gol.curr = gol.next //swap
		for i := range tmp {
			for j := range tmp[i] {
				tmp[i][j] = DeadCell
			}
		}
		gol.next = tmp
		gol.printCurrentBoard()
	}
}

//set the live cell on the board
func (gol *Life) setCell(row, col int) *Life {
	//ignore+log invalid input for cell
	if row >= gol.size || col >= gol.size || row < 0 || col < 0 {
		fmt.Println("Invalid row/col: row:", row, "col: ", col, " size: ", gol.size)
		return gol
	}
	gol.curr[row][col] = LiveCell
	return gol
}

// for test/demo only
func initialize(size int) *Life {
	gol := &Life{}
	boardA := make(board, size)
	boardB := make(board, size)
	for i := 0; i < size; i++ {
		boardA[i] = make([]int, size)
		boardB[i] = make([]int, size)
	}
	gol.curr = boardA
	gol.next = boardB
	gol.size = size
	return gol
}

func main() {
	// glider demo with 25 x 25 board
	initialize(25).
		setCell(10, 11).
		setCell(12, 10).
		setCell(12, 11).
		setCell(12, 12).
		setCell(11, 12).
		evolve(250 * time.Millisecond)
}
