package main

import (
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet"
)

func ReadExcel(filename string, c Config) (spreadsheet.Row, []StudentRecord) {
	fmt.Printf("Reading in file: %s\n", filename)
	ss, err := spreadsheet.Open(filename)
	handleErr(err, "Error opening spreadsheet")
	defer ss.Close()
	sheet := ss.Sheets()[0]
	records := []StudentRecord{}
	headerRow := sheet.Row(uint32(c.HeaderStartsOnRow))
	fmt.Printf("Header row %d\n", c.HeaderStartsOnRow)
	for idx, r := range sheet.Rows() {
		rowNumber := idx + 1
		// idx is 0-based, but rows in excel are 1-based
		// So +1 so that rowNumber is actually the row number
		if rowNumber <= c.HeaderStartsOnRow {
			fmt.Printf("Skipping row %d - second cell in skipped row is %s\n", rowNumber, r.Cell(fmt.Sprintf("B")).GetString())
			continue
		}
		record := NewStudentRecordFromRow(r, c)
		fmt.Printf("Reading row %d: %s %s\n", rowNumber, record.FirstName, record.LastName)
		if record.FirstName == "" {
			fmt.Println("FirstName blank - returning")
			return headerRow, records
		}
		records = append(records, record)
	}
	return headerRow, records
}

func RowToMap(r spreadsheet.Row) map[string]string {
	m := make(map[string]string)
	for _, c := range r.Cells() {
		column, err := c.Column()
		handleErr(err, "Error reading column")
		m[column] = c.GetString()
	}
	return m
}
