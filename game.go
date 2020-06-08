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
var directions = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

func (gol *Life) printCurrentBoard() {
	fmt.Print("\n")
	for i := range gol.curr {
		for j := range gol.curr[i] {
			if gol.curr[i][j] == LiveCell {
				fmt.Print("^ ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Print("\n")
	}
}

func (gol *Life) applyLifeRules(row, col int) {
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

func (gol *Life) resetNext() {
	gol.next = make(board, gol.size)
	for i := range gol.next {
		gol.next[i] = make([]int, gol.size)
	}
}

func (gol *Life) evolve() {
    if gol.size < 1 {
        fmt.Print("size of board cannot be less than 1")
        return
    }
	gol.printCurrentBoard()
	for range time.Tick(333 * time.Millisecond) {
		for i := range gol.curr {
			for j := range gol.curr[i] {
				gol.applyLifeRules(i, j)
			}
		}
		gol.curr = gol.next //swap
		gol.resetNext()     //reset
		gol.printCurrentBoard()
	}
}

// for test/demo only
func initWithGlider() *Life {
    gol := &Life{}
    size := 25
    boardA := make(board, size)
    boardB := make(board, size)
    for i := range boardA {
        boardA[i] = make([]int, size)
        boardB[i] = make([]int, size)
    }
    //glider pattern
    boardA[10][11] = LiveCell
    boardA[12][10] = LiveCell
    boardA[12][11] = LiveCell
    boardA[12][12] = LiveCell
    boardA[11][12] = LiveCell
    gol.curr = boardA
    gol.next = boardB
    gol.size = size
    return gol
}

func main() {
    // glider demo with 25 x 25 board
	initWithGlider().evolve()
}
