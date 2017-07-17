package modules

import "time"

var (
	FileLoc = "C:/Program Files/dstr"
	// FileLoc    = "C:/go/bin/src/JPRO/tmp"
	MAX        = ReadTxtMax()
	TimeFormat = "2006-01-02 15:04:05"
	StartTime  = time.Now().Format(TimeFormat)
	EndTime    = time.Now().Format(TimeFormat)
	TempMaxs   []ExcelTable
	AllTempMax []ExcelTable
)
