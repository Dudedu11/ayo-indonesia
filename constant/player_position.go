package constant

const (
	PositionPenyerang     = "penyerang"
	PositionGelandang     = "gelandang"
	PositionBertahan      = "bertahan"
	PositionPenjagaGawang = "penjaga gawang"
)

var ValidPositions = []string{
	PositionPenyerang,
	PositionGelandang,
	PositionBertahan,
	PositionPenjagaGawang,
}

func IsValidPosition(pos string) bool {
	for _, p := range ValidPositions {
		if p == pos {
			return true
		}
	}
	return false
}
