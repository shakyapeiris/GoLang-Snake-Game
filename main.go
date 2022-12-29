package main

import (
	"fmt"
	"strings"
)

func main() {
	reset()
	moves := make(chan string, 1)
	ended := make(chan string, 1)

	select {
	case msg := <-ended:
		fmt.Println(msg)
		return
	default:
		go updateSnake(ended)
		go listenDirnChanges(moves)

		ended <- "Ouch!"
		close(ended)
		for {
			var command string

			fmt.Println("Please enter a command: ")
			_, err := fmt.Scan(&command)
			if err != nil {
				fmt.Println(err)
			}
			command = strings.ToLower(command)

			switch command {
			case "move":
				var newDirn string
				fmt.Print("Enter a direction: ")
				_, err := fmt.Scan(&newDirn)
				if err != nil {
					fmt.Println(err)
				}
				moves <- newDirn
				break
			case "print":
				printBoard()
				break
			case "snake":
				fmt.Println(snake)
				break
			case "direction":
				getCurrentDirn()
				break
			case "apple":
				fmt.Println(appleCoordinates)
				break
			case "score":
				fmt.Println(score)
				break
			default:
				fmt.Println("Invalid command!")
			}
		}
	}

}
