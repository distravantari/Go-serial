package main

import (
	"fmt"
	"io/ioutil"

	"github.com/tealeg/xlsx"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteTxt(text string) {

	d1 := []byte(text)
	err := ioutil.WriteFile(FileLoc+"/Raw.txt", d1, 0644)
	check(err)
}

func WriteRep(start string, end string, typ string) { // typ means the type (log/txt/excel)

	var temp string

	data, _ := ioutil.ReadFile(FileLoc + "/Rep.log")

	temp += string(data) + "\n"
	if typ == "log" {
		temp += FileLoc + `/Rep.log
		` + start + ` Data: start
		` + end + ` Data: end`
	}

	d1 := []byte(temp)
	err := ioutil.WriteFile(FileLoc+"/Rep.log", d1, 0644)
	check(err)
}

func WriteMax(acts []ExcelTable) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	xFiles, err := xlsx.OpenFile(FileLoc + "/Max.xlsx")
	if err != nil {
		file = xlsx.NewFile()
		sheet, err = file.AddSheet("Sheet1")
		if err != nil {
			fmt.Printf(err.Error())
		}
		row = sheet.AddRow()
		TableHeader(row, cell) // add table header
		createTable(sheet, row, cell, acts)
		err = file.Save(FileLoc + "/Max.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
	} else {
		file = xFiles
		sheet = file.Sheets[0]
		createTable(sheet, row, cell, acts)
		err = file.Save(FileLoc + "/Max.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

}

func createTable(sheet *xlsx.Sheet, row *xlsx.Row, cell *xlsx.Cell, acts []ExcelTable) {

	for i := range acts {
		row = sheet.AddRow()
		TableBody(row, cell, acts[i])
	}

}
