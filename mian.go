package main

import (
	"net/http"

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
	c.JSON(http.StatusOK, exampleCrossword1)
}

var exampleCrossword1 CrosswordView = CrosswordView{
	SizeHorizontal: 20,
	SizeVertical:   20,
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
