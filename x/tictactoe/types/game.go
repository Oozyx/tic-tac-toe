package types

func InitializeBoard() []BoardEntry {
	board := []BoardEntry{
		BoardEntry_EMPTY, BoardEntry_EMPTY, BoardEntry_EMPTY,
		BoardEntry_EMPTY, BoardEntry_EMPTY, BoardEntry_EMPTY,
		BoardEntry_EMPTY, BoardEntry_EMPTY, BoardEntry_EMPTY,
	}

	return board
}

func GetWinMatrix() [][]uint {
	winMatrix := [][]uint{
		{0, 1, 2}, // Check first row.
		{3, 4, 5}, // Check second Row
		{6, 7, 8}, // Check third Row
		{0, 3, 6}, // Check first column
		{1, 4, 7}, // Check second Column
		{2, 5, 8}, // Check third Column
		{0, 4, 8}, // Check first Diagonal
		{2, 4, 6}, // Check second Diagonal
	}

	return winMatrix
}

func (game Game) IsWinner(player BoardEntry) bool {
	winMatrix := GetWinMatrix()
	for i := 0; i < 8; i++ {
		if game.Board[winMatrix[i][0]] == player &&
			game.Board[winMatrix[i][1]] == player &&
			game.Board[winMatrix[i][2]] == player {
			return true
		}
	}

	return false
}

func (game Game) IsGameComplete() bool {
	for i := 0; i < 9; i++ {
		if game.Board[i] == BoardEntry_EMPTY {
			return false
		}
	}

	return true
}