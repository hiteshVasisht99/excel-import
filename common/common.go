package common

import (
	"errors"

	"github.com/hiteshVasisht99/excel-import/models"
	"github.com/xuri/excelize/v2"
)

// ParseExcelFile parses the Excel file and returns records
func ParseExcelFile(filePath string) ([]*models.Employee, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	// Retrieve the names of all sheets
	sheets := f.GetSheetList()
	if len(sheets) != 1 {
		return nil, errors.New("expected exactly one sheet in the Excel file")
	}
	sheet := sheets[0]
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	var records []*models.Employee
	for i, row := range rows {
		// Skip header row
		if i == 0 {
			continue
		}

		if len(row) < 10 {
			return nil, errors.New("row has insufficient columns")
		}

		employee := &models.Employee{
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			Country:     row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		}

		records = append(records, employee)
	}

	return records, nil
}
