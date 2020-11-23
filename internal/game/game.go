package game

import "strconv"

// Game is a state of a NoNoNet game
type Game struct {
	Mask board
	Tries board
	HelpersX [10][5]int
	HelpersY [10][5]int
	PosX int
	PosY int
	Errors int
}

func (g *Game) Select() string {
	if g.Mask[g.PosY][g.PosX] {
		g.Tries[g.PosY][g.PosX] = true
	} else {
		g.Errors++ 
	}
	return g.Render()
}

func (g *Game) MoveUp() string {
	if g.PosY > 0 {
		g.PosY--
	}
	return g.Render()
}

func (g *Game) MoveDown() string {
	if g.PosY < 9 {
		g.PosY++ 
	}
	return g.Render()
}

func (g *Game) MoveRight() string {
	if g.PosX < 9 {
		g.PosX++
	}
	return g.Render()
}

func (g *Game) MoveLeft() string {
	if g.PosX > 0 {
		g.PosX--
	}
	return g.Render()
}

func (g *Game) Render() string {
	var buffer [20][40]string
	for i, a := range buffer {
		for j := range a {
			buffer[i][j] = " "
		}
	}
	// Draw x Helpers
	for i, a := range g.HelpersX {
		col := (i + 7) * 2
		for j, b := range a {
			if b > 0 {
				buffer[j + 1][col] = strconv.Itoa(b)
			}
		}
	}

	// Draw y Helpers
	for i, a := range g.HelpersY {
		col := i + 7
		for j, b := range a {
			if b > 0 {
				buffer[col][(j + 1) * 2] = strconv.Itoa(b)
			}
		}
	}

	// Draw tries
	for y, line := range g.Tries {
		for x, cell := range line {
			if cell {
				buffer[y + 7][(x + 7) * 2] = "\u001b[47m \u001b[0m"
				buffer[y + 7][(x + 7) * 2 + 1] = "\u001b[47m \u001b[0m"
			}
		}
	}

	// Draw cursor
	var cursor string
	if g.Tries[g.PosY][g.PosX] {
		cursor = "\u001b[47m\u001b[30;1mX\u001b[0m"
	} else {
		cursor = "X"
	}
	buffer[g.PosY + 7][(g.PosX + 7) * 2] = cursor
	buffer[g.PosY + 7][((g.PosX + 7) * 2) + 1] = cursor

	// Draw errors
	if g.Errors >= 1 {
		buffer[5][4] = "X"
		if g.Errors >= 2 {
			buffer[5][5] = "X"
			if g.Errors >= 3 {
				buffer[5][3] = "L"
				buffer[5][4] = "O"
				buffer[5][5] = "S"
				buffer[5][6] = "T"
			}
		}		
	}

	var won = true
	for y, line := range g.Tries {
		for x, cell := range line {
			won = won && (g.Mask[y][x] == cell)
		}
	}
	if won {
		buffer[5][3] = "W"
		buffer[5][4] = "O"
		buffer[5][5] = "N"
		buffer[5][6] = "!"
	}

	var output string
	for _, line := range buffer {
		for _, cell := range line {
			output += cell
		}
		output = (output + "\r\n") 
	}
	return output
}

var InitTries board = board{
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false},
}

var CoffeeBoard board = board{
	{false, false, true, false, true, false, true, false, false, false},
	{false, false, true, false, true, false, true, false, false, false},
	{false, false, false, false, false, false, false, false, false, false },
	{false, true, true, true, true, true, true, true, false, false },
	{false, true, true, false ,true, true, true, true, true, true },
	{false, true, true, false ,true, true, true, true, false, true },
	{false, true, true, true, true, true, true, true, true, false },
	{false, true, true, true, true, true, true, true, false, false },
	{true, false, true, true, true, true, true, false, false, true},
	{false, true, true, true, true, true, true, true, true, false},
}

var CoffeeBoardHelpersX [10][5]int = [10][5]int{
	{0, 0, 0, 0, 1},
	{0, 0, 0, 5, 1},
	{0, 0, 0, 2, 7},
	{0, 0, 0, 1, 4},
	{0, 0, 0, 2, 7},
	{0, 0, 0, 0, 7},
	{0, 0, 0, 2, 7},
	{0, 0, 0, 5, 1},
	{0, 0, 1, 1, 1},
	{0, 0, 0, 3, 1},
}

var CoffeeBoardHelpersY [10][5]int = [10][5]int{
	{0, 0, 1, 1, 1},
	{0, 0, 1, 1, 1},
	{0, 0, 0, 0, 0},
	{0, 0, 0, 0, 7},
	{0, 0, 0, 2, 6},
	{0, 0, 2, 4, 1},
	{0, 0, 0, 0, 8},
	{0, 0, 0, 0, 7},
	{0, 0, 1, 5, 1},
	{0, 0, 0, 0, 8},
}