package main

import (
	"testing"
)

func TestLiveCellLivesWithExactly2or3Nbrs(t *testing.T) {
	board := initialize(25).
		setCell(3, 6).
		setCell(4, 6).
		setCell(5, 6)
	board.applyLifeRules(4, 6)
	if board.countLiveNbrs(4, 6) != 2 {
		t.Error("bad test expecting exactly 2 neighbors")
	}
	if board.next[4][6] == DeadCell {
		t.Error("cell with exactly 2 live neighbors stays alive")
	}
	board = initialize(25).
		setCell(3, 6).
		setCell(4, 6).
		setCell(4, 7).
		setCell(5, 6)
	board.applyLifeRules(4, 6)
	if board.countLiveNbrs(4, 6) != 3 {
		t.Error("bad test expecting exactly 3 neighbors")
	}
	if board.next[4][6] == DeadCell {
		t.Error("cell with exactly 3 live neighbors stays alive")
	}
}
func TestDeadCellLivesWithExactly3LiveNgbrs(t *testing.T) {
	board := initialize(25).
		setCell(3, 6).
		setCell(4, 6).
		setCell(5, 6).
		setCell(3, 7).
		setCell(5, 7).
		setCell(3, 8).
		setCell(4, 8)
	board.applyLifeRules(4, 8)
	if board.next[4][8] == DeadCell {
		t.Error("cell with exactly 3 live neighbors comes alive")
	}
}

func TestCellDeathRuleLessThan2Nbrs(t *testing.T) {
	board := initialize(25).
		setCell(3, 6).
		setCell(4, 6)
	board.applyLifeRules(3, 6)
	board.applyLifeRules(4, 6)
	if board.next[4][6] == LiveCell || board.next[3][6] == LiveCell {
		t.Error("cell with less than two neighbors must ")
	}
}

func TestCellDeathRuleMoreThan3Nbrs(t *testing.T) {
	board := initialize(25).
		setCell(3, 6).
		setCell(4, 6).
		setCell(5, 6).
		setCell(3, 7).
		setCell(4, 7).
		setCell(5, 7).
		setCell(3, 8).
		setCell(4, 8).
		setCell(5, 8)
	board.applyLifeRules(4, 7)
	if board.next[4][7] == LiveCell && board.next[3][6] == DeadCell {
		t.Error("cell with more than 3 neighbors must die, others live")
	}
}
func TestLiveNbors(t *testing.T) {
	board := initialize(25).
		setCell(3, 6).
		setCell(4, 6).
		setCell(5, 6).
		setCell(3, 7).
		setCell(4, 7).
		setCell(5, 7).
		setCell(3, 8).
		setCell(4, 8).
		setCell(5, 8)

	nbrs := board.countLiveNbrs(4, 7)
	if nbrs != 8 {
		t.Error("incorrect neighbors")
	}
	nbrs = board.countLiveNbrs(3, 6)
	if nbrs != 3 {
		t.Error("incorrect neighbors")
	}

	nbrs = board.countLiveNbrs(23, 23)
	if nbrs != 0 {
		t.Error("incorrect neighbors")
	}
}
