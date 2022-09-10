package crossword_generator

type WordReference interface {
	*Word | *TableWord
	GetLetter(i int) rune
	String() string
}

func (word *Word) GetLetter(i int) rune {
	return word.Value[i]
}

func (word *TableWord) GetLetter(i int) rune {
	return word.Value[i]
}
