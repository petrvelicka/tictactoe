package main

import (
	"errors"
	"fmt"
	"github.com/inancgumus/screen"
	"math/rand"
)

var computer = []bool{false, false, false}
var symbols = []string{" ", "✕", "○"}
var lines = []string{"A", "B", "C"}

// printField prints game field
func printField(field [][]int) {
	fmt.Println(" |1|2|3")
	for i := 0; i < 3; i++ {
		fmt.Println("-|-|-|-")
		fmt.Printf("%s|%s|%s|%s\n", lines[i], symbols[field[i][0]], symbols[field[i][1]], symbols[field[i][2]])
	}
}

// checkSame check if the array has all elements equal and any of them isn't zero
func checkSame(line []int) bool {
	product := 1
	for _, elem := range line {
		product *= elem
	}

	if product == 0 {
		return false
	}

	val := line[0]
	for _, elem := range line[1:] {
		if elem != val {
			return false
		}
	}

	return true
}

// checkWin returns 0 if no player won and there is available move
// It returns -1 if there is a draw and a number of player if they won
func checkWin(field [][]int) int {
	won := -1

	// draw check
	for _, line := range field {
		for _, elem := range line {
			if elem == 0 {
				won = 0
			}
		}
	}

	for _, line := range field {
		if checkSame(line) {
			won = line[0]
		}
	}

	for i := range field[0] {
		arr := []int{}
		for j := 0; j < len(field); j += 1 {
			arr = append(arr, field[j][i])
		}

		if checkSame(arr) {
			won = arr[0]
		}
	}

	diag := []int{}
	side := []int{}
	for i := range field {
		diag = append(diag, field[i][i])
		side = append(side, field[i][len(field)-i-1])
	}

	if checkSame(diag) {
		won = diag[0]
	}

	if checkSame(side) {
		won = side[0]
	}
	return won
}

// getInput gets input from a user and returns coordinates of a field to play or an error
func getInput() ([]int, error) {
	var x string
	var y int

	if _, err := fmt.Scanf("%s %d", &x, &y); err != nil {
		return []int{0, 0}, err
	}

	if y < 1 || y > 3 {
		return []int{0, 0}, errors.New("column number invalid")
	}

	result := []int{0, y - 1}

	matched := false
	for i, value := range lines {
		if x == value {
			result[0] = i
			matched = true
		}
	}
	if !matched {
		return []int{0, 0}, errors.New("row number invalid")
	}

	return result, nil
}

// play is the game loop
func play() {
	field := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	plays := 0
	for {
		screen.Clear()
		screen.MoveTopLeft()
		plays += 1
		if plays > 2 {
			plays = 1
		}
		fmt.Printf("Now plays player %v\n", plays)
		printField(field)

		if !computer[plays] {
			for {
				input, err := getInput()
				if err != nil {
					fmt.Println("Error: " + err.Error())
				} else {
					if field[input[0]][input[1]] == 0 {
						field[input[0]][input[1]] = plays
						break
					} else {
						fmt.Println("Error: field is not available")
					}
				}
			}
		} else {
			for {
				x := rand.Intn(3)
				y := rand.Intn(3)

				if field[x][y] == 0 {
					field[x][y] = plays
					break
				}
			}
		}

		result := checkWin(field)
		if result > 0 {
			fmt.Printf("Player number %v won\n", result)
			break
		} else if result == -1 {
			fmt.Println("Draw")
			break
		}
	}
}

func welcome() {
	fmt.Println("Welcome to TicTacToe")
	fmt.Print("Do you want to play with a computer (Y/n)? ")
	var line string
	_, _ = fmt.Scanln(&line)
	if line == "y" {
		computer[2] = true
	}
}

func main() {
	welcome()
	play()
}
