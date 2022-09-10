package main

import (
	"net/http"

	"crossword-service/crossword_generator"

	"github.com/gin-gonic/gin"
)

func main() {

	handler := CreateHandler()

	router := gin.Default()
	router.POST("/crosswords", handler.CreateCrossword)

	router.Run()
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type Handler struct {
}

func CreateHandler() *Handler {
	h := Handler{}

	return &h
}

// func (h *Handler) CreateCrossword(c *gin.Context) {
// 	c.JSON(http.StatusOK,  map[string]interface{}{
// 		 "Crossword": h.ExampleCrossword,
// 	})
// }

func (h *Handler) CreateCrossword(c *gin.Context) {

	// stringWords := [...]string{"sator", "arepo", "tenet", "opera", "rotas"}

	// words := make([]*Word, 0)
	// for _, stringWord := range stringWords {
	// 	words = append(words, NewWord(stringWord))
	// }

	var words []*crossword_generator.Word = crossword_generator.CrissCross27{}.GetWords()

	var resultingTables []*crossword_generator.WordsTable = crossword_generator.MainLoop(words)

	var horizontalWords []PositionedWordView = make([]PositionedWordView, 0)
	var verticalWords []PositionedWordView = make([]PositionedWordView, 0)

	for _, word := range resultingTables[0].GetHorizontalWords() {
		horizontalWords = append(horizontalWords, *CreatePositionedWordViewFromTableWord(word))
	}

	for _, word := range resultingTables[0].GetVerticalWords() {
		verticalWords = append(verticalWords, *CreatePositionedWordViewFromTableWord(word))
	}

	var exampleCrossword2 CrosswordView = CrosswordView{
		SizeHorizontal:  crossword_generator.HorizontalSize,
		SizeVertical:    crossword_generator.VerticalSize,
		HorizontalWords: horizontalWords,
		VerticalWords:   verticalWords,
	}

	c.JSON(http.StatusOK, exampleCrossword2)
}

var exampleCrossword1 CrosswordView = CrosswordView{
	SizeHorizontal: crossword_generator.HorizontalSize,
	SizeVertical:   crossword_generator.VerticalSize,
	HorizontalWords: []PositionedWordView{
		*(word("sator").CreatePositionedWordView(10, 10)),
		*CreatePositionedWordView("tenet", 10, 12),
		*CreatePositionedWordView("rotas", 10, 14),
	},
	VerticalWords: []PositionedWordView{
		*CreatePositionedWordView("arepo", 11, 10),
		*CreatePositionedWordView("opera", 13, 10),
	},
}
