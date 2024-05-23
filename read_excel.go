package main

import (
	"fmt"

	"github.com/unidoc/unioffice/spreadsheet"
)

func ReadExcel(filename string, c Config) (spreadsheet.Row, []StudentRecord) {
	fmt.Printf("Reading in file: %s\n", filename)
	ss, err := spreadsheet.Open(filename)
	if err != nil {
		panic(err)
	}
	defer ss.Close()
	sheet := ss.Sheets()[0]
	records := []StudentRecord{}
	headerRow := sheet.Row(uint32(c.HeaderStartsOnRow))
	for rowNumber, r := range sheet.Rows() {
		if rowNumber <= c.HeaderStartsOnRow {
			continue
		}
		record := NewStudentRecordFromRow(r, c)
		if record.FirstName == "" {
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
		if err != nil {
			panic(err)
		}
		m[column] = c.GetString()
	}
	return m
}
