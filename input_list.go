package main

import crs "crossword-service/crossword_generator"

type InputList struct {
	Words []string `binding:"required"`
}

func (il *InputList) GetWords() []*crs.Word {
	words := make([]*crs.Word, 0)

	for _, aString := range il.Words {
		aWord := crs.NewWord(aString)
		words = append(words, aWord)
	}

	return words
}
