package crosswd

type TableWord struct {
	*Word
	X                   int
	Y                   int
	Direction           WordDirection
	BusyCharacters      []bool
	InnerTablePositions []*TablePosition
}

func NewTableWord(value string, x int, y int, direction WordDirection, innerTablePositions []*TablePosition) *TableWord {
	newTableWord := new(TableWord)
	newTableWord.Word = NewWord(value)
	newTableWord.X = x
	newTableWord.Y = y
	newTableWord.Direction = direction
	newTableWord.BusyCharacters = make([]bool, len(value))
	newTableWord.InnerTablePositions = innerTablePositions

	return newTableWord
}

func NewTableWordFromWord(word *Word, x int, y int, direction WordDirection, innerTablePositions []*TablePosition) *TableWord {
	newTableWord := new(TableWord)
	newTableWord.Word = word
	newTableWord.X = x
	newTableWord.Y = y
	newTableWord.Direction = direction
	newTableWord.BusyCharacters = make([]bool, len(word.Value))
	newTableWord.InnerTablePositions = innerTablePositions

	return newTableWord
}

func CloneTableWord(oldTableWord *TableWord) *TableWord {
	newTableWord := new(TableWord)
	newTableWord.Word = CloneWord(oldTableWord.Word)
	newTableWord.X = oldTableWord.X
	newTableWord.Y = oldTableWord.Y
	newTableWord.Direction = oldTableWord.Direction
	newTableWord.BusyCharacters = make([]bool, len(oldTableWord.Value))

	copy(newTableWord.BusyCharacters, oldTableWord.BusyCharacters)

	return newTableWord
}

func (w *TableWord) String() string {
	return string(w.Value)
}
