package main

import (
	"net/http"

	crs "crossword-service/crossword_generator"

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

	var words []*crs.Word = crs.CrissCross27{}.GetWords()

	var resultingTables []*crs.WordsTable = crs.MainLoop(words)

	var horizontalWords []PositionedWordView = make([]PositionedWordView, 0)
	var verticalWords []PositionedWordView = make([]PositionedWordView, 0)

	for _, word := range resultingTables[0].GetHorizontalWords() {
		horizontalWords = append(horizontalWords, MakePositionedWordViewFromTableWord(word))
	}

	for _, word := range resultingTables[0].GetVerticalWords() {
		valueOfTypeTabword := tabword(*word)
		verticalWords = append(verticalWords, (&valueOfTypeTabword).MakePositionedWordViewFromTableWord())
	}

	var exampleCrossword2 CrosswordView = CrosswordView{
		SizeHorizontal:  crs.HorizontalSize,
		SizeVertical:    crs.VerticalSize,
		HorizontalWords: horizontalWords,
		VerticalWords:   verticalWords,
	}

	c.JSON(http.StatusOK, exampleCrossword2)
}

var exampleCrossword1 CrosswordView = CrosswordView{
	SizeHorizontal: crs.HorizontalSize,
	SizeVertical:   crs.VerticalSize,
	HorizontalWords: []PositionedWordView{
		*(word("sator").CreatePositionedWordView(10, 10)),
		MakePositionedWordView("tenet", 10, 12),
		MakePositionedWordView("rotas", 10, 14),
	},
	VerticalWords: []PositionedWordView{
		MakePositionedWordView("arepo", 11, 10),
		MakePositionedWordView("opera", 13, 10),
	},
}
