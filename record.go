package main

import (
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet"
)

type StudentRecord struct {
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
		if err != nil {
			panic(err)
		}
		results[i] = int(val)
	}

	nc, err := dataRow.Cell(config.NumberCorrect).GetValueAsNumber()
	if err != nil {
		panic(err)
	}
	pc, err := dataRow.Cell(config.PercentCorrect).GetValueAsNumber()
	if err != nil {
		panic(err)
	}
	nia, err := dataRow.Cell(config.NumberItemsAttempted).GetValueAsNumber()
	if err != nil {
		panic(err)
	}

	return StudentRecord{
		LastName:             dataRow.Cell(config.LastName).GetString(),
		FirstName:            dataRow.Cell(config.FirstName).GetString(),
		NumberCorrect:        int(nc),
		PercentCorrect:       pc,
		NumberItemsAttempted: int(nia),
		Results:              results,
	}
}

func NewStudentRecordFromRow1(row spreadsheet.Row, startColumn string, nResults int) StudentRecord {
	results := make([]int, nResults)

	for i := 0; i < nResults; i++ {
		col := GetColumnFrom(startColumn, i)
		fmt.Printf("Getting column %s\n", col)
		val, err := row.Cell(col).GetValueAsNumber()
		if err != nil {
			panic(err)
		}
		results[i] = int(val)
	}
	fmt.Printf("hi\n")
	nc, err := row.Cell("E").GetValueAsNumber()
	if err != nil {
		panic(err)
	}
	pc, err := row.Cell("F").GetValueAsNumber()
	if err != nil {
		panic(err)
	}
	nia, err := row.Cell("G").GetValueAsNumber()
	if err != nil {
		panic(err)
	}
	return StudentRecord{
		LastName:             row.Cell("B").GetString(),
		FirstName:            row.Cell("C").GetString(),
		NumberCorrect:        int(nc),
		PercentCorrect:       pc,
		NumberItemsAttempted: int(nia),
		Results:              results,
	}
}
