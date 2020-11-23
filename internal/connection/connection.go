package connection

import (
	"fmt"
	"firlus.dev/nononet/internal/game"
	"net"
	"bufio"
	escapes "github.com/snugfox/ansi-escapes"
)

// HandleConnection pro forma
func HandleConnection(conn net.Conn) {
	game := game.Game{
		game.CoffeeBoard,
		game.InitTries,
		game.CoffeeBoardHelpersX,
		game.CoffeeBoardHelpersY,
		0, 0, 0}
	buffer := bufio.NewReader(conn)
	conn.Write([]byte("\377\375\042\377\373\001"))
	conn.Write([]byte(escapes.ClearScreen + game.Render()))
	for {
		var output string
		command, _ := buffer.ReadByte()
		fmt.Println("new command")
		if command == 27 {
			command, _ = buffer.ReadByte()
			if command == 91 {
				command, _ = buffer.ReadByte()
				switch command {
				case 65: // up arrow
					output = game.MoveUp()
				case 66: // arrow down
					output = game.MoveDown()
				case 67: // right arrow
					output = game.MoveRight()
				case 68: // left arrow
					output = game.MoveLeft()
				}
			}
		} else if command == 32 {
			output = game.Select()
		}
		if output == "" {
			output = game.Render()
		}
		conn.Write([]byte(escapes.ClearScreen + string(output)))
	}
}

func generateOutput(pos int) string {
	output := ""
	for i := 0; i < 100; i++ {
		if i == pos {
			output += "X"
		} else {
			output += " "
		}
		if i % 10 == 9 {
			output += "\r\n"
		}
	}
	return output
}