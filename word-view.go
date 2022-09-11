package main

import (
	crosswd "crossword-service/crossword_generator"
	"strings"
)

type PositionedWordView struct {
	Value string `json:"w"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

func MakePositionedWordView(value string, x int, y int) PositionedWordView {
	w := PositionedWordView{}
	w.Value = strings.ToUpper(value)
	w.X = x
	w.Y = y

	return w
}

type word string

func (s word) CreatePositionedWordView(x int, y int) *PositionedWordView {
	w := new(PositionedWordView)
	w.Value = strings.ToUpper(string(s))
	w.X = x
	w.Y = y

	return w
}

func MakePositionedWordViewFromTableWord(wr *crosswd.TableWord) PositionedWordView {
	w := PositionedWordView{}
	w.Value = string(wr.Word.Value)
	w.X = wr.X
	w.Y = wr.Y

	return w
}

type tabword crosswd.TableWord

func (wr *tabword) MakePositionedWordViewFromTableWord() PositionedWordView {
	w := PositionedWordView{}
	w.Value = string(wr.Word.Value)
	w.X = wr.X
	w.Y = wr.Y

	return w
}
