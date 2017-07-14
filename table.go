package main

import "github.com/tealeg/xlsx"

func TableHeader(row *xlsx.Row, cell *xlsx.Cell) {
	cell = row.AddCell()
	cell.Value = "No"
	cell = row.AddCell()
	cell.Value = "Jam"
	cell = row.AddCell()
	cell.Value = "Max"
	cell = row.AddCell()
	cell.Value = "Lama"
	cell = row.AddCell()
	cell.Value = "Awal"
	cell = row.AddCell()
	cell.Value = "Akhir"
	// cell = row.AddCell()
	// cell.Value = ""
	// cell = row.AddCell()
	// cell.Value = ""
	// cell = row.AddCell()
	// cell.Value = "Tanggal"
	// cell = row.AddCell()
	// cell.Value = "Unknown"
	// cell = row.AddCell()
	// cell.Value = "Unknown"
	// cell = row.AddCell()
	// cell.Value = "Aktifitas"
}

type ExcelTable struct {
	No    string
	Jam   string
	Max   string
	Lama  string
	Awal  string
	Akhir string
	// Tanggal   string
	// Unknown   string
	// Unknown   string
	// Aktifitas string
}

func TableBody(row *xlsx.Row, cell *xlsx.Cell, act ExcelTable) {
	cell = row.AddCell()
	cell.Value = act.No
	cell = row.AddCell()
	cell.Value = act.Jam
	cell = row.AddCell()
	cell.Value = act.Max
	cell = row.AddCell()
	cell.Value = act.Lama
	cell = row.AddCell()
	cell.Value = act.Awal
	cell = row.AddCell()
	cell.Value = act.Akhir
	// cell = row.AddCell()
	// cell.Value = ""
	// cell = row.AddCell()
	// cell.Value = ""
	// cell = row.AddCell()
	// cell.Value = "Tanggal"
	// cell = row.AddCell()
	// cell.Value = "Unknown"
	// cell = row.AddCell()
	// cell.Value = "Unknown"
	// cell = row.AddCell()
	// cell.Value = "Aktifitas"
}
