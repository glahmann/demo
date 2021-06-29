package board

// func Create() [][]string {

// }

func ToString(board [][]string) (boardStr string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			boardStr += board[i][j]
			if j < len(board[0])-1 {
				boardStr += " , "
			}
		}
		boardStr += "\n"
	}
	return
}
