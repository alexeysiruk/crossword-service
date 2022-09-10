package crossword_generator

import (
	"fmt"
	"sort"
	"strconv"
)

const WordTablesCountLimit int = 5000

func MainLoop(words []*Word) []*WordsTable {
	wordsTable := NewWordsTable()

	firstWord := words[0]
	words = words[1:]

	wordsTable.AddWordHorizontally(firstWord, (HorizontalSize-len(firstWord.Value))/2, (VerticalSize-1)/2)

	nextStepTables := make([]*WordsTable, 0)
	currentStepTables := make([]*WordsTable, 0)

	currentStepTables = append(currentStepTables, wordsTable)

	currentDirection := Horizontal

	wordsAddedCounter := 1

	for len(words) > 0 {
		if len(nextStepTables) > 0 {
			currentStepTables = make([]*WordsTable, len(nextStepTables))
			copy(currentStepTables, nextStepTables)
			nextStepTables = make([]*WordsTable, 0)
		}
		wordToAdd := words[0]
		words = words[1:]

		fmt.Println("Words added count:", wordsAddedCounter)
		fmt.Println("Word to add: " + wordToAdd.String())

		currentDirection = currentDirection.Inverted()

		for _, table := range currentStepTables {
			tables := WordAdditionStep(table, wordToAdd, currentDirection)
			nextStepTables = append(nextStepTables, tables...)
		}

		fmt.Println("Next step tables count:", len(nextStepTables))

		if len(nextStepTables) > 0 {
			wordsAddedCounter++

			// fmt.Println("Tables before sorting: ")
			// for _, t := range nextStepTables {
			// 	fmt.Print(t.crossingsCounter, ", ")
			// }
			// fmt.Println()

			sort.Slice(nextStepTables, func(i, j int) bool {
				return nextStepTables[i].crossingsCounter > nextStepTables[j].crossingsCounter
			})

			// fmt.Println("Tables after sorting: ")
			// for _, t := range nextStepTables {
			// 	fmt.Print(t.crossingsCounter, ", ")
			// }
			// fmt.Println()

			if len(nextStepTables) > WordTablesCountLimit {
				nextStepTables = nextStepTables[:WordTablesCountLimit]
			}
		}
	}

	if len(nextStepTables) == 0 {
		nextStepTables = currentStepTables
	}

	// fmt.Println("Final words added count:", wordsAddedCounter)
	// for _, aTable := range nextStepTables {
	// 	fmt.Println(aTable)
	// }

	return nextStepTables
}

func WordAdditionStep(wordsTable *WordsTable, wordToAdd *Word, wordDirection WordDirection) []*WordsTable {
	var availableTableCrossingPoints []*CrossingPoint[*TableWord] = wordsTable.GetAvailableCrossingPointsFromTableDirection(wordDirection.Inverted())
	//fmt.Println(availableTableCrossingPoints)

	var wordToAddCrossingPoints []*CrossingPoint[*Word] = GetWordToAddCrossingPoints(wordToAdd)
	//fmt.Println(wordToAddCrossingPoints)

	var potentialAdditions = GetPotentialCrossings(availableTableCrossingPoints, wordToAddCrossingPoints)

	var resultingTables = make([]*WordsTable, 0)
	var additionUniquenessSet = make(map[string]bool)

	for _, potentialAddition := range potentialAdditions {
		//fmt.Println(potentialAddition.String() + " " + potentialAddition.Peer.String())
		InitializePositions(potentialAddition)
		var additionStartCoordiantesHash string = strconv.Itoa(potentialAddition.NewWordCache.StartX) + ":" + strconv.Itoa(potentialAddition.NewWordCache.StartY)

		if _, ok := additionUniquenessSet[additionStartCoordiantesHash]; !ok {
			isAdditionFenced := wordsTable.IsAdditionFenced(potentialAddition)
			if !isAdditionFenced {
				continue
			}

			isAdditionCompatible := wordsTable.IsAdditionCompatible(potentialAddition)
			if !isAdditionCompatible {
				continue
			}

			isAdditionIsolated := wordsTable.IsAdditionIsolated(potentialAddition)
			if !isAdditionIsolated {
				continue
			}

			var newWordsTable = CloneWordsTable(wordsTable)

			newWordsTable.AddWordFromPotentialCrossing(potentialAddition)
			resultingTables = append(resultingTables, newWordsTable)
			//fmt.Println(newWordsTable)

			additionUniquenessSet[additionStartCoordiantesHash] = true
		}
	}

	// for _, aResultingTable := range resultingTables {
	// 	fmt.Println(aResultingTable)
	// }

	return resultingTables
}
