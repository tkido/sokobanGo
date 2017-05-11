package main

// General Consts
const (
	GOAL = 1 << iota
	MAN
	AVOID
	BAG
	WALL
	EXTRA
	BLANK
	FLOOR = 0

	MONG = MAN | GOAL  // MAN on GOAL
	MONA = MAN | AVOID // MAN on AVOID
	BONG = BAG | GOAL  // BAG on GOAL
)
