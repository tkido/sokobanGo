package main

// Level of sokoban screen
type Level struct {
	Name        string
	Description string
	Man         int
}

func parse(path string) {
	readMap := map[string]int{
		" ": FLOOR,
		".": GOAL,
		"@": MAN,
		"+": MONG,
		"$": BAG,
		"*": BONG,
		"#": WALL,
	}

}
