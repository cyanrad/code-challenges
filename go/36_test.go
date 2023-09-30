package leet

import (
	"fmt"
	"log"
	"testing"
)

func resetTempCount(list []int) {
	for i := 0; i < 10; i++ {
		list[i] = 0
	}
}

func isValidSudoku(board [][]byte) bool {
	tempCount := make([]int, 10)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			if tempCount[board[i][j]-48] == 1 {
				fmt.Println("here 1")
				return false
			}

			tempCount[board[i][j]-48]++
		}

		resetTempCount(tempCount)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}

			if tempCount[board[j][i]-48] == 1 {
				fmt.Println("here 2")
				return false
			}

			tempCount[board[j][i]-48]++
		}

		resetTempCount(tempCount)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			index_x := j%3 + (i/3)*3
			index_y := (j / 3) + (i%3)*3

			log.Println(index_x, index_y)
			if board[index_x][index_y] == '.' {
				continue
			}

			if tempCount[board[index_x][index_y]-48] == 1 {
				// return false
			}

			tempCount[board[index_x][index_y]-48]++
		}

		log.Println("resetting")

		resetTempCount(tempCount)
	}

	return true
}

func TestIsValidSudoku(t *testing.T) {
	board := [][]byte{{'.', '.', '.', '.', '5', '.', '.', '1', '.'}, {'.', '4', '.', '3', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '3', '.', '.', '1'}, {'8', '.', '.', '.', '.', '.', '.', '2', '.'}, {'.', '.', '2', '.', '7', '.', '.', '.', '.'}, {'.', '1', '5', '.', '.', '.', '.', '.', '.'}, {'.', '.', '.', '.', '.', '2', '.', '.', '.'}, {'.', '2', '.', '9', '.', '.', '.', '.', '.'}, {'.', '.', '4', '.', '.', '.', '.', '.', '.'}}
	val := isValidSudoku(board)
	t.Log(val)
	if val != false {
		t.Fail()
	}
}
