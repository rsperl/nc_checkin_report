package main

import (
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet"
)

type StudentRecord struct {
	Core                 string
	LastName             string
	FirstName            string
	NumberCorrect        int
	PercentCorrect       float64
	NumberItemsAttempted int
	Results              []int
}

var arr = [...]string{".", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func GetColumnFrom(startColumn string, n int) string {
	if n == 0 {
		return startColumn
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == startColumn {
			if i+n > len(arr) {
				panic("column out of range")
			}
			return arr[i+n]
		}

	}
	return ""
}

func NewStudentRecordFromRow(dataRow spreadsheet.Row, config Config) StudentRecord {
	results := make([]int, config.Results.Count)
	for i := 0; i < config.Results.Count; i++ {
		col := IncrementLetter(config.Results.StartsOnColumn, i)
		val, err := dataRow.Cell(col).GetValueAsNumber()
		handleErr(err, fmt.Sprintf("Error reading results from column %s", col))
		results[i] = int(val)
	}

	nc, err := dataRow.Cell(config.NumberCorrect).GetValueAsNumber()
	handleErr(err, fmt.Sprintf("Error reading number correct from column %s", config.NumberCorrect))

	pc, err := dataRow.Cell(config.PercentCorrect).GetValueAsNumber()
	handleErr(err, fmt.Sprintf("Error reading percent correct from column %s", config.PercentCorrect))

	nia, err := dataRow.Cell(config.NumberItemsAttempted).GetValueAsNumber()
	handleErr(err, fmt.Sprintf("Error reading number items attempted from column %s", config.NumberItemsAttempted))

	return StudentRecord{
		Core:                 dataRow.Cell(config.Core).GetString(),
		LastName:             dataRow.Cell(config.LastName).GetString(),
		FirstName:            dataRow.Cell(config.FirstName).GetString(),
		NumberCorrect:        int(nc),
		PercentCorrect:       pc,
		NumberItemsAttempted: int(nia),
		Results:              results,
	}
}
