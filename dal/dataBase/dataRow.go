package dataBase

import (
	"errors"
)

type DataRow struct {
	table *DataTable

	cells []interface{}
}

func (this *DataRow) Cells() []interface{} {
	return this.cells
}

func (this *DataRow) Len() int {
	return len(this.cells)
}

func (this *DataRow) Cell(celIndex int) (interface{}, error) {
	if len(this.cells) <= celIndex {
		return nil, errors.New("cell out of range")
	}

	return this.cells[celIndex], nil
}

func (this *DataRow) CellByName(cellName string) (interface{}, error) {
	celIndex := this.table.cellIndex(cellName)
	if celIndex < 0 {
		return nil, errors.New("cell name no exist")
	}

	return this.cells[celIndex], nil
}

func newDataRow(_table *DataTable, _cells []interface{}) *DataRow {
	return &DataRow{
		table: _table,
		cells: _cells,
	}
}
