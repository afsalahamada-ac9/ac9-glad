/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package util

import (
	"strconv"
	"strings"
	"ac9/glad/pkg/common"
	"time"
)

type ValueExtractor = func(int) []interface{}

func DBTimeNow() string {
	return time.Now().UTC().Format(common.DBFormatDateTimeMS)
}

// GenBulkInsertPGSQL generates bulk insert statement for postgres
func GenBulkInsertPGSQL(
	tableName string,
	columns []string,
	numRows int,
	valueExtractor ValueExtractor,
) (string, []interface{}) {
	numCols := len(columns)
	var queryBuilder strings.Builder
	queryBuilder.WriteString("INSERT INTO ")
	queryBuilder.WriteString(tableName)
	queryBuilder.WriteString("(")
	for i, column := range columns {
		queryBuilder.WriteString("\"")
		queryBuilder.WriteString(column)
		queryBuilder.WriteString("\"")
		if i < numCols-1 {
			queryBuilder.WriteString(",")
		}
	}
	queryBuilder.WriteString(") VALUES ")
	var valueArgs []interface{}
	valueArgs = make([]interface{}, 0, numRows*numCols)
	for rowIndex := 0; rowIndex < numRows; rowIndex++ {
		queryBuilder.WriteString("(")
		for colIndex := 0; colIndex < numCols; colIndex++ {
			queryBuilder.WriteString("$")
			queryBuilder.WriteString(strconv.Itoa(rowIndex*numCols + colIndex + 1))
			if colIndex < numCols-1 {
				queryBuilder.WriteString(",")
			}
		}
		queryBuilder.WriteString(")")
		if rowIndex < numRows-1 {
			queryBuilder.WriteString(",")
		}
		valueArgs = append(valueArgs, valueExtractor(rowIndex)...)
	}
	return queryBuilder.String(), valueArgs
}

// GenBulkDeletePGSQL generates bulk delete statement for postgres
func GenBulkDeletePGSQL(
	tableName string,
	columns []string,
	numRows int,
	valueExtractor ValueExtractor,
) (string, []interface{}) {
	numCols := len(columns)
	var queryBuilder strings.Builder
	queryBuilder.WriteString("DELETE FROM ")
	queryBuilder.WriteString(tableName)
	queryBuilder.WriteString(" WHERE ")

	var valueArgs []interface{}
	valueArgs = make([]interface{}, 0, numRows*numCols)
	for rowIndex := 0; rowIndex < numRows; rowIndex++ {
		if rowIndex > 0 {
			queryBuilder.WriteString(" OR (")
		} else {
			queryBuilder.WriteString("(")
		}
		for colIndex := 0; colIndex < numCols; colIndex++ {
			queryBuilder.WriteString(columns[colIndex] + " = ")
			queryBuilder.WriteString("$")
			queryBuilder.WriteString(strconv.Itoa(rowIndex*numCols + colIndex + 1))
			if colIndex < numCols-1 {
				queryBuilder.WriteString(" AND ")
			}
		}
		queryBuilder.WriteString(")")
		if rowIndex == numRows-1 {
			queryBuilder.WriteString(";")
		}
		valueArgs = append(valueArgs, valueExtractor(rowIndex)...)
	}
	return queryBuilder.String(), valueArgs
}
