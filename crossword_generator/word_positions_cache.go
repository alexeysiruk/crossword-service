package crosswd

type WordPositionsCache struct {
	StartX              int
	StartY              int
	Direction           WordDirection
	InnerTablePositions []*TablePosition
}

func InitializePositions(potentialAddition *PotentialCrossing[*Word, *TableWord]) {
	newWordPositionsCache := new(WordPositionsCache)
	newWordPositionsCache.InnerTablePositions = make([]*TablePosition, len(potentialAddition.Word.Value))

	if potentialAddition.Peer.Word.Direction == Horizontal {
		newWordPositionsCache.StartX = potentialAddition.Peer.Word.X + potentialAddition.Peer.Index
		newWordPositionsCache.StartY = potentialAddition.Peer.Word.Y - potentialAddition.Index

		for i := 0; i < len(potentialAddition.Word.Value); i++ {
			newWordPositionsCache.InnerTablePositions[i] = new(TablePosition)
			newWordPositionsCache.InnerTablePositions[i].X = newWordPositionsCache.StartX
			newWordPositionsCache.InnerTablePositions[i].Y = newWordPositionsCache.StartY + i
		}

		newWordPositionsCache.Direction = Vertical
	} else {
		newWordPositionsCache.StartX = potentialAddition.Peer.Word.X - potentialAddition.Index
		newWordPositionsCache.StartY = potentialAddition.Peer.Word.Y + potentialAddition.Peer.Index

		for i := 0; i < len(potentialAddition.Word.Value); i++ {
			newWordPositionsCache.InnerTablePositions[i] = new(TablePosition)
			newWordPositionsCache.InnerTablePositions[i].X = newWordPositionsCache.StartX + i
			newWordPositionsCache.InnerTablePositions[i].Y = newWordPositionsCache.StartY
		}

		newWordPositionsCache.Direction = Horizontal
	}

	potentialAddition.NewWordCache = newWordPositionsCache
}
