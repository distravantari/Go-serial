package modules

import "time"

var (
	FileLoc = "C:/tools/system" + time.Now().Format("2006-01-02")
	// FileLoc               = "C:/go/bin/src/JPRO/tmp/JPRO" + time.Now().Format("2006-01-02")
	MAX                   = ReadTxtMax()
	TimeFormat            = "2006-01-02 15:04:05"
	StartTime             = time.Now().Format(TimeFormat)
	EndTime               = time.Now().Format(TimeFormat)
	TempMaxs              []ExcelTable
	AllTempMax            []ExcelTable
	directWriteAllTempMax []ExcelTable
)
