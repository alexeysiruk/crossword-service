package crossword_generator

import "strings"

type Word struct {
	Value []rune
}

func NewWord(value string) *Word {
	w := new(Word)
	w.Value = []rune(strings.ToUpper(value))

	return w
}

func CloneWord(oldWord *Word) *Word {
	w := new(Word)
	w.Value = oldWord.Value

	return w
}

func (w *Word) String() string {
	return string(w.Value)
}
