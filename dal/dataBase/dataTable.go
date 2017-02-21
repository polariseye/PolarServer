package dataBase

import (
	"database/sql"
	"errors"
)

type DataTable struct {
	rowData     []*DataRow
	columnNames map[string]int
}

func (this *DataTable) init(rows *sql.Rows) error {
	defer func() {
		rows.Close()
	}()

	// 读取列信息和保存列名称
	tmpColumns, errMsg := rows.Columns()
	if errMsg != nil {
		return errMsg
	}
	for index, val := range tmpColumns {
		this.columnNames[val] = index
	}

	// 读取行数据
	this.rowData = make([]*DataRow, 0)
	columnCount := len(this.columnNames)
	for rows.Next() {
		rowCells := make([]interface{}, columnCount)
		rows.Scan(rowCells...)

		this.rowData = append(this.rowData, newDataRow(this, rowCells))
	}

	return nil
}

func (this *DataTable) CellByIndex(rowIndex int, cellIndex int) (interface{}, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	rowItem := this.rowData[rowIndex]
	if len(rowItem.cells) <= cellIndex {
		return nil, errors.New("column out of range")
	}

	return rowItem.cells[cellIndex], nil
}

func (this *DataTable) CellByCellName(rowIndex int, cellName string) (interface{}, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	cellIndex, isExist := this.columnNames[cellName]
	if isExist == false {
		return nil, errors.New("column no Exist")
	}

	rowItem := this.rowData[rowIndex]
	return rowItem.cells[cellIndex], nil
}

func (this *DataTable) Row(rowIndex int) (*DataRow, error) {
	if len(this.rowData) <= rowIndex {
		return nil, errors.New("row out of range")
	}

	return this.rowData[rowIndex], nil
}

func (this *DataTable) cellIndex(cellName string) int {
	cellIndex, isExist := this.columnNames[cellName]
	if isExist == false {
		return -1
	}

	return cellIndex
}

func (this *DataTable) Len() int {
	return len(this.rowData)
}

func NewDataTable(rows *sql.Rows) (*DataTable, error) {
	table := new(DataTable)
	errMsg := table.init(rows)
	if errMsg != nil {
		return nil, errMsg
	}

	return table, nil
}
