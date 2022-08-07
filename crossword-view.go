package main

type CrosswordView struct {
	SizeHorizontal  int                  `json:"width"`
	SizeVertical    int                  `json:"height"`
	HorizontalWords []PositionedWordView `json:"horizontals"`
	VerticalWords   []PositionedWordView `json:"verticals"`
}
