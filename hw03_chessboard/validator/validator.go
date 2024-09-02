package validator

const (
	MaxHeightChessboard = 50
	MaxWidthChessboard  = 50
)

func ValidateHeightChessboard(height int) bool {
	return validateValueChessboard(height, MaxHeightChessboard)
}

func ValidateWidthChessboard(height int) bool {
	return validateValueChessboard(height, MaxWidthChessboard)
}

func validateValueChessboard(value int, limit int) bool {
	if value < 1 || value > limit {
		return false
	}

	return true
}
