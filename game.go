package main

import (
	"fmt"
	"math/rand"
	"time"
)

const boardSize = 20

var board = make([]string, boardSize)
var snake [][2]int
var dirn string
var appleCoordinates [2]int
var score int

func reset() {
	board = make([]string, boardSize)
	snake = append(snake, [2]int{5, 5})
	dirn = "D"
	appleCoordinates = [2]int{rand.Int() % boardSize, rand.Int() % boardSize}
	score = 0
}

// updateSnake will update snake's location according to it's moving direction
func updateSnake(ended chan string) {
	time.Sleep(1 * time.Second)
	var newSnake [][2]int
	snakeHead := snake[0]

	switch dirn {
	case "U":
		if snakeHead[1] == 0 {
			snakeHead[1] = 20
		}
		snakeHead = [2]int{snakeHead[0], snakeHead[1] - 1}
		break
	case "D":
		if snakeHead[1] == 19 {
			snakeHead[1] = -1
		}
		snakeHead = [2]int{snakeHead[0], snakeHead[1] + 1}
		break
	case "L":
		if snakeHead[0] == 0 {
			snakeHead[0] = 20
		}
		snakeHead = [2]int{snakeHead[0] - 1, snakeHead[1]}
		break
	case "R":
		if snakeHead[0] == 19 {
			snakeHead[0] = -1
		}
		snakeHead = [2]int{snakeHead[0] + 1, snakeHead[1]}
		break
	}
	newSnake = [][2]int{snakeHead}
	length := len(snake) - 1
	if snakeHead == appleCoordinates {
		score++
		length++
		appleCoordinates = [2]int{rand.Int() % boardSize, rand.Int() % boardSize}
	}
	for i := 0; i < length; i++ {
		newSnake = append(newSnake, snake[i])
		if snake[i] == snakeHead {
			ended <- "Ouch!"
			close(ended)
		}
	}

	snake = newSnake
	updateSnake(ended)
}

// listenDirnChanges update current moving direction
func listenDirnChanges(moves chan string) {
	dirn = <-moves
}

// printBoard prints current board
func printBoard() {
	for i := 0; i < boardSize; i++ {
		fmt.Print("|")
		for j := 0; j < boardSize; j++ {

			fmt.Print(" |")
		}
		fmt.Println()
	}
	fmt.Println()
}

func getCurrentDirn() {
	switch dirn {
	case "U":
		fmt.Println("Upwards")
		break
	case "D":
		fmt.Println("Downwards")
		break
	case "L":
		fmt.Println("Leftwards")
		break
	case "R":
		fmt.Println("Rightwards")
		break
	}
}
