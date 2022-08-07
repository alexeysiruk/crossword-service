package main

import (
	"strings"
)

type PositionedWordView struct {
	Value string `json:"w"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

func CreatePositionedWordView(value string, x int, y int) *PositionedWordView {
	w := new(PositionedWordView)
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
