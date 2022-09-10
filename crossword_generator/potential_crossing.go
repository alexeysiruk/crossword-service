package crossword_generator

type PotentialCrossing[T WordReference, U WordReference] struct {
	*CrossingPoint[T]
	Peer         *CrossingPoint[U]
	NewWordCache *WordPositionsCache
}

func NewPotentialCrossing(wordToAddCrossingPoint *CrossingPoint[*Word], tableCrossingPoint *CrossingPoint[*TableWord]) *PotentialCrossing[*Word, *TableWord] {
	newPotentialCrossing := new(PotentialCrossing[*Word, *TableWord])
	newPotentialCrossing.CrossingPoint = wordToAddCrossingPoint
	newPotentialCrossing.Peer = tableCrossingPoint

	return newPotentialCrossing
}
