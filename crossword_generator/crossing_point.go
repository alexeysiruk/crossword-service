package crossword_generator

import "strconv"

type CrossingPoint[T WordReference] struct {
	Word   T
	Index  int
	Letter rune
}

func NewCrossingPoint[T WordReference](word T, i int) *CrossingPoint[T] {
	newCrossingPoint := new(CrossingPoint[T])
	newCrossingPoint.Word = word
	newCrossingPoint.Index = i
	newCrossingPoint.Letter = word.GetLetter(i)

	return newCrossingPoint
}

func (p *CrossingPoint[T]) String() string {
	return p.Word.String() + ":" + strconv.Itoa(p.Index) + ":" + string(p.Letter)
}
