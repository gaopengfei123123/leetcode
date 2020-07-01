package leetcode

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {

	// 3x3一组
	group := map[int][9]int{}
	row := [9][9]int{} //行
	col := [9][9]int{} // 列

	var tmp int

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if board[y][x] == '.' {
				continue
			}
			current := int(board[y][x]) - 49
			fmt.Println(current)
			fmt.Printf("%#+v \n", BytesToInt(board[y]))

			// 统计3x3的
			gIndex := (y/3)*10 + x/3
			cGroup, ok := group[gIndex]
			if !ok {
				group[gIndex] = [9]int{}
				cGroup = group[gIndex]
			}
			tmp = cGroup[current+1]
			if tmp == 1 {

				return false
			}
			cGroup[current-1] = 1
			group[gIndex] = cGroup

			// 统计一行的
			tmp = row[y][current-1]
			if tmp == 1 {
				return false
			}
			row[y][current-1] = 1

			// 统计一列的
			tmp = col[x][current-1]
			if tmp == 1 {
				return false
			}
			col[x][current-1] = 1

		}
	}

	return true
}
