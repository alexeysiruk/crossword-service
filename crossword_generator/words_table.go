package crosswd

import (
	"reflect"
	"strconv"
	"strings"
)

const HorizontalSize int = 20
const VerticalSize int = 20

type WordsTable struct {
	innerTable      [HorizontalSize][VerticalSize]rune
	tableWords      []*TableWord
	horizontalWords []*TableWord
	verticalWords   []*TableWord

	horizontalsCache          [HorizontalSize][VerticalSize]rune
	verticalsCache            [HorizontalSize][VerticalSize]rune
	horizontalsIsolationCache [HorizontalSize][VerticalSize]bool
	verticalsIsolationCache   [HorizontalSize][VerticalSize]bool

	crossingsCounter int
}

func NewWordsTable() *WordsTable {
	newWordsTable := new(WordsTable)
	newWordsTable.tableWords = make([]*TableWord, 0)
	newWordsTable.horizontalWords = make([]*TableWord, 0)
	newWordsTable.verticalWords = make([]*TableWord, 0)

	return newWordsTable
}

func CloneWordsTable(oldWordsTable *WordsTable) *WordsTable {
	newWordsTable := new(WordsTable)
	newWordsTable.innerTable = oldWordsTable.innerTable //supposed to copy

	newWordsTable.tableWords = make([]*TableWord, 0)
	for _, w := range oldWordsTable.tableWords {
		newWordsTable.tableWords = append(newWordsTable.tableWords, CloneTableWord(w))
	}

	newWordsTable.horizontalWords = make([]*TableWord, 0)
	newWordsTable.verticalWords = make([]*TableWord, 0)
	for _, w := range newWordsTable.tableWords {
		if w.Direction == Horizontal {
			newWordsTable.horizontalWords = append(newWordsTable.horizontalWords, w)
		} else {
			newWordsTable.verticalWords = append(newWordsTable.verticalWords, w)
		}
	}

	newWordsTable.horizontalsCache = oldWordsTable.horizontalsCache
	newWordsTable.verticalsCache = oldWordsTable.verticalsCache
	newWordsTable.horizontalsIsolationCache = oldWordsTable.horizontalsIsolationCache
	newWordsTable.verticalsIsolationCache = oldWordsTable.verticalsIsolationCache

	newWordsTable.crossingsCounter = oldWordsTable.crossingsCounter

	return newWordsTable
}

func (t *WordsTable) GetHorizontalWords() []*TableWord {
	result := make([]*TableWord, 0)

	for _, tableWord := range t.horizontalWords {
		result = append(result, tableWord)
	}

	return result
}

func (t *WordsTable) GetVerticalWords() []*TableWord {
	result := make([]*TableWord, 0)

	for _, tableWord := range t.verticalWords {
		result = append(result, tableWord)
	}

	return result
}

func (t *WordsTable) String() string {
	var sb strings.Builder

	for i := 0; i < VerticalSize; i++ {
		needToPrintRow := false
		for j := 0; j < HorizontalSize; j++ {
			if t.innerTable[j][i] != 0 { //||
				// 	t.horizontalsCache[j][i] != 0 ||
				// 	t.verticalsCache[j][i] != 0 ||
				// 	t.horizontalsIsolationCache[j][i] ||
				// 	t.verticalsIsolationCache[j][i] {
				needToPrintRow = true
				break
			}
		}
		if needToPrintRow {
			for j := 0; j < HorizontalSize; j++ {
				if t.innerTable[j][i] != 0 {
					sb.WriteRune(t.innerTable[j][i])
				} else {
					sb.WriteString("_")
				}
			}
			// sb.WriteString("|")
			// for j := 0; j < HorizontalSize; j++ {
			// 	if t.horizontalsCache[j][i] != 0 {
			// 		sb.WriteRune(t.horizontalsCache[j][i])
			// 	} else {
			// 		sb.WriteString("_")
			// 	}
			// }
			// sb.WriteString("|")
			// for j := 0; j < HorizontalSize; j++ {
			// 	if t.verticalsCache[j][i] != 0 {
			// 		sb.WriteRune(t.verticalsCache[j][i])
			// 	} else {
			// 		sb.WriteString("_")
			// 	}
			// }
			// sb.WriteString("|")
			// for j := 0; j < HorizontalSize; j++ {
			// 	if t.horizontalsIsolationCache[j][i] {
			// 		sb.WriteString("#")
			// 	} else {
			// 		sb.WriteString("_")
			// 	}

			// }
			// sb.WriteString("|")
			// for j := 0; j < HorizontalSize; j++ {
			// 	if t.verticalsIsolationCache[j][i] {
			// 		sb.WriteString("#")
			// 	} else {
			// 		sb.WriteString("_")
			// 	}
			// }

			sb.WriteString("\r\n")
		}
	}
	sb.WriteString("Crossings: " + strconv.Itoa(t.crossingsCounter) + "\r\n")

	return sb.String()
}

func (wt *WordsTable) AddWord(newWord *Word, newWordStartX int, newWordStartY int, newWordDirection WordDirection, newWordTablePositions []*TablePosition) *TableWord {
	//Caching the pre- and post- isolation zones
	firstPosition := newWordTablePositions[0]
	lastPosition := newWordTablePositions[len(newWordTablePositions)-1]

	if newWordDirection == Horizontal {
		if firstPosition.X > 0 {
			wt.horizontalsIsolationCache[firstPosition.X-1][firstPosition.Y] = true
		}
		if lastPosition.X < HorizontalSize-1 {
			wt.horizontalsIsolationCache[lastPosition.X+1][lastPosition.Y] = true
		}
	} else {
		if firstPosition.Y > 0 {
			wt.verticalsIsolationCache[firstPosition.X][firstPosition.Y-1] = true
		}
		if lastPosition.Y < VerticalSize-1 {
			wt.verticalsIsolationCache[lastPosition.X][lastPosition.Y+1] = true
		}
	}
	//end Caching the pre - and post - isolation zones

	for i := 0; i < int(len(newWord.Value)); i++ {
		var c = newWordTablePositions[i]
		var currentValue = wt.innerTable[c.X][c.Y]
		var valueToAdd = newWord.Value[i]
		if currentValue != 0 && currentValue != valueToAdd {
			panic("Wrong letter is being added to the crossword!")
		}

		wt.innerTable[c.X][c.Y] = valueToAdd

		//Caching
		if newWordDirection == Horizontal {
			wt.horizontalsCache[c.X][c.Y] = valueToAdd
			wt.horizontalsIsolationCache[c.X][c.Y] = true
			if c.Y > 0 {
				wt.horizontalsIsolationCache[c.X][c.Y-1] = true
			}
			if c.Y < VerticalSize-1 {
				wt.horizontalsIsolationCache[c.X][c.Y+1] = true
			}
			//Counting the crossings
			if wt.verticalsCache[c.X][c.Y] != 0 {
				wt.crossingsCounter++
			}
		} else {
			wt.verticalsCache[c.X][c.Y] = valueToAdd
			wt.verticalsIsolationCache[c.X][c.Y] = true
			if c.X > 0 {
				wt.verticalsIsolationCache[c.X-1][c.Y] = true
			}
			if c.X < HorizontalSize-1 {
				wt.verticalsIsolationCache[c.X+1][c.Y] = true
			}
			//Counting the crosisngs
			if wt.horizontalsCache[c.X][c.Y] != 0 {
				wt.crossingsCounter++
			}
		}
		//end Caching
	}

	var newTableWord = NewTableWordFromWord(newWord, newWordStartX, newWordStartY, newWordDirection, newWordTablePositions)
	wt.tableWords = append(wt.tableWords, newTableWord)
	if newWordDirection == Horizontal {
		wt.horizontalWords = append(wt.horizontalWords, newTableWord)
	} else {
		wt.verticalWords = append(wt.verticalWords, newTableWord)
	}

	return newTableWord
}

func (wt *WordsTable) AddWordHorizontally(newWord *Word, newWordStartX int, newWordStartY int) *TableWord {
	var newWordTablePositions []*TablePosition = make([]*TablePosition, len(newWord.Value))

	for i := 0; i < len(newWord.Value); i++ {
		newWordTablePositions[i] = &TablePosition{newWordStartX + i, newWordStartY}
	}

	return wt.AddWord(newWord, newWordStartX, newWordStartY, Horizontal, newWordTablePositions)
}

func (wt *WordsTable) AddWordVertically(newWord *Word, newWordStartX int, newWordStartY int) *TableWord {
	var newWordTablePositions []*TablePosition = make([]*TablePosition, len(newWord.Value))

	for i := 0; i < len(newWord.Value); i++ {
		newWordTablePositions[i] = &TablePosition{newWordStartX, newWordStartY + i}
	}

	return wt.AddWord(newWord, newWordStartX, newWordStartY, Vertical, newWordTablePositions)
}

func (wt *WordsTable) AddWordFromPotentialCrossing(crossingToAdd *PotentialCrossing[*Word, *TableWord]) {
	if crossingToAdd.Peer.Word.BusyCharacters[crossingToAdd.Peer.Index] {
		panic("existing table word has the potential crossing point busy")
	}

	newTableWord := wt.AddWord(
		crossingToAdd.Word,
		crossingToAdd.NewWordCache.StartX,
		crossingToAdd.NewWordCache.StartY,
		crossingToAdd.NewWordCache.Direction,
		crossingToAdd.NewWordCache.InnerTablePositions)
	newTableWord.BusyCharacters[crossingToAdd.Index] = true

	clonedPeerTableWord := wt.GetTableWordForClonesSynchronization(crossingToAdd.Peer.Word)
	clonedPeerTableWord.BusyCharacters[crossingToAdd.Peer.Index] = true
}

func (wt *WordsTable) GetTableWordForClonesSynchronization(originalTableWord *TableWord) *TableWord {
	tableWords := wt.tableWords

	for _, tableWord := range tableWords {
		if (tableWord.X == originalTableWord.X) &&
			(tableWord.Y == originalTableWord.Y) &&
			(tableWord.Direction == originalTableWord.Direction) &&
			(reflect.DeepEqual(tableWord.Value, originalTableWord.Value)) {
			return tableWord
		}
	}

	panic("Cannot find the cloned table word")
}

func GetAvailableCrossingPointsFromTableWords(filteredWords *[]*TableWord) []*CrossingPoint[*TableWord] {
	result := make([]*CrossingPoint[*TableWord], 0)

	for _, tableWord := range *filteredWords {
		for index := range tableWord.Value {
			if !tableWord.BusyCharacters[index] {
				newCrossingPoint := NewCrossingPoint(tableWord, index)
				result = append(result, newCrossingPoint)
			}
		}
	}

	return result
}

func (wt *WordsTable) GetAvailableCrossingPointsFromTableDirection(direction WordDirection) []*CrossingPoint[*TableWord] {
	var result []*CrossingPoint[*TableWord]

	if direction == Horizontal {
		result = GetAvailableCrossingPointsFromTableWords(&wt.horizontalWords)
	} else {
		result = GetAvailableCrossingPointsFromTableWords(&wt.verticalWords)
	}

	return result
}

func GetWordToAddCrossingPoints(wordToAdd *Word) []*CrossingPoint[*Word] {
	result := make([]*CrossingPoint[*Word], 0)

	for index := range wordToAdd.Value {
		newCrossingPoint := NewCrossingPoint(wordToAdd, index)
		result = append(result, newCrossingPoint)
	}

	return result
}

func GetPotentialCrossings(availableTableCrossingPoints []*CrossingPoint[*TableWord], wordToAddCrossingPoints []*CrossingPoint[*Word]) []*PotentialCrossing[*Word, *TableWord] {
	result := make([]*PotentialCrossing[*Word, *TableWord], 0)

	for i := 0; i < len(availableTableCrossingPoints); i++ {
		for j := 0; j < len(wordToAddCrossingPoints); j++ {
			if availableTableCrossingPoints[i].Letter == wordToAddCrossingPoints[j].Letter {
				newPotentialCrossing := NewPotentialCrossing(wordToAddCrossingPoints[j], availableTableCrossingPoints[i])
				result = append(result, newPotentialCrossing)
			}
		}
	}

	return result
}

func (wt *WordsTable) IsAdditionFenced(potentialAddition *PotentialCrossing[*Word, *TableWord]) bool {
	var isFenced bool = true

	positions := potentialAddition.NewWordCache.InnerTablePositions

	if positions[0].X < 0 ||
		positions[0].Y < 0 ||
		positions[len(positions)-1].X > HorizontalSize-1 ||
		positions[len(positions)-1].Y > VerticalSize-1 {
		isFenced = false
	}

	return isFenced
}

func (wt *WordsTable) IsAdditionCompatible(potentialAddition *PotentialCrossing[*Word, *TableWord]) bool {
	var isCompatible bool = true

	for i, currentPosition := range potentialAddition.NewWordCache.InnerTablePositions {
		currentValue := wt.innerTable[currentPosition.X][currentPosition.Y]

		if currentValue != 0 && currentValue != potentialAddition.Word.Value[i] {
			isCompatible = false
			break
		}
	}

	return isCompatible
}

func (wt *WordsTable) IsAdditionIsolated(potentialAddition *PotentialCrossing[*Word, *TableWord]) bool {
	var positions []*TablePosition = potentialAddition.NewWordCache.InnerTablePositions

	var first = positions[0]
	var last = positions[len(positions)-1]

	if potentialAddition.NewWordCache.Direction == Horizontal {
		//there is a place before the word
		if first.X > 0 {
			//joint vertical at start
			if wt.verticalsCache[first.X-1][first.Y] != 0 {
				return false
			}
			//horizontal joint pre
			if wt.horizontalsCache[first.X-1][first.Y] != 0 {
				return false
			}
		}
		//there is a place after the word
		if last.X < HorizontalSize-1 {
			//joint vertical at end
			if wt.verticalsCache[last.X+1][last.Y] != 0 {
				return false
			}
			//horizontal joint post
			if wt.horizontalsCache[last.X+1][last.Y] != 0 {
				return false
			}
		}

		for i := 0; i < len(positions); i++ {
			c := positions[i]
			if wt.horizontalsIsolationCache[c.X][c.Y] {
				//horizontal overlapping
				if wt.horizontalsCache[c.X][c.Y] != 0 {
					return false
					//a position above or below a horizontal word which is not taken by a vertical word
				} else if wt.verticalsCache[c.X][c.Y] == 0 {
					return false
				}
			}
			//joint vertical below/above
			if wt.verticalsCache[c.X][c.Y] == 0 && ((c.Y < VerticalSize-1 && wt.verticalsCache[c.X][c.Y+1] != 0) ||
				(c.Y > 0 && wt.verticalsCache[c.X][c.Y-1] != 0)) {
				return false
			}
		}
	} else {
		//there is a place before the word
		if first.Y > 0 {
			//joint horizontal above at start
			if wt.horizontalsCache[first.X][first.Y-1] != 0 {
				return false
			}
			//vertical joint pre-
			if wt.verticalsCache[first.X][first.Y-1] != 0 {
				return false
			}
		}
		//there is a place after the word
		if last.Y < VerticalSize-1 {
			//joint horizontal below at end
			if wt.horizontalsCache[last.X][last.Y+1] != 0 {
				return false
			}
			//vertical joint post-
			if wt.verticalsCache[last.X][last.Y+1] != 0 {
				return false
			}
		}

		for i := 0; i < len(positions); i++ {
			c := positions[i]
			if wt.verticalsIsolationCache[c.X][c.Y] {
				//vertical overlapping
				if wt.verticalsCache[c.X][c.Y] != 0 {
					return false
					//a position left or right a vertical word which is not taken by a horizontal word
				} else if wt.horizontalsCache[c.X][c.Y] == 0 {
					return false
				}
			}
			//joint horizontal on left/right
			if wt.horizontalsCache[c.X][c.Y] == 0 && ((c.X < HorizontalSize-1 && wt.horizontalsCache[c.X+1][c.Y] != 0) ||
				(c.X > 0 && wt.horizontalsCache[c.X-1][c.Y] != 0)) {
				return false
			}
		}
	}

	return true
}
