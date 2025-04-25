package main

import (
	"fmt"

	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

// Use the golang library github.com/unidoc/unioffice to create a function that
// writes a new word document. Each page should have a table that shows
// StudentRecord data.

func populateCell(row *document.Row, label string, value string) {
	cell := row.AddCell()
	cell.Properties().Margins().SetLeft(5 * measurement.Point)
	paraLabel := cell.AddParagraph()
	paraLabel.Properties().Spacing().SetBefore(2 * measurement.Point)

	cell = row.AddCell()
	cell.Properties().Margins().SetLeft(5 * measurement.Point)
	paraValue := cell.AddParagraph()
	paraLabel.AddRun().AddText(label)
	paraValue.AddRun().AddText(value)
}

func WriteDoc(filename string, records []StudentRecord, headers map[string]string, config Config) error {
	doc := document.New()
	defer doc.Close()

	// Iterate over your StudentRecord data and add a table on each page
	for _, record := range records {
		para := doc.AddParagraph()
		para.SetStyle("Heading2")
		run := para.AddRun()
		run.AddText(config.Title)
		table := doc.AddTable()
		table.Properties().SetWidthPercent(50)
		borders := table.Properties().Borders()
		borders.SetAll(wml.ST_BorderSingle, color.Auto, 1*measurement.Point)

		row := table.AddRow()
		populateCell(&row, headers[config.Core], record.Core)

		row = table.AddRow()
		populateCell(&row, headers[config.FirstName], record.FirstName)

		row = table.AddRow()
		populateCell(&row, headers[config.LastName], record.LastName)

		row = table.AddRow()
		value := fmt.Sprintf("%d/%d", record.NumberCorrect, config.PointsPossible)
		populateCell(&row, headers[config.NumberCorrect], value)

		row = table.AddRow()
		pc := fmt.Sprintf("%.0f", record.PercentCorrect)
		populateCell(&row, headers[config.PercentCorrect], pc)

		for i, r := range record.Results {
			row = table.AddRow()
			column := IncrementLetter(config.Results.StartsOnColumn, i)
			label := headers[column]
			populateCell(&row, label, fmt.Sprintf("%d", r))
		}

		doc.AddParagraph().AddRun().AddPageBreak()
	}

	// Save the document to a file
	fmt.Printf("Writing to file: %s\n", filename)
	return doc.SaveToFile(filename)
}

// Write a function that accepts a letter and an integer and returns the letter incremented by the integer. For example, "A", 1 -> "B", "A", 2 -> "C", "Z", 1 -> "AA", "Z", 2 -> "AB", etc.
func IncrementLetter(letter string, n int) string {
	if len(letter) == 0 {
		return ""
	}
	return string(rune(letter[0]) + rune(n))
}
