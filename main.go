package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	FileLoc    = "C:/Users/Desktop/go/src/JPRO/tmp"
	MAX        = 200
	TimeFormat = "2006-01-02 15:04:05"
	StartTime  = time.Now().Format(TimeFormat)
	EndTime    = time.Now().Format(TimeFormat)
	TempMaxs   []ExcelTable
	AllTempMax []ExcelTable
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

func OnClose() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	WriteRep(StartTime, EndTime, "log")
	if len(AllTempMax) > 0 {
		WriteMax(AllTempMax)
	}
	fmt.Println("exiting")
	os.Exit(3)
}

func ConfigRuntime() {
	go OnClose()
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	var t *testing.T
	go Counter(t)
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")

	router.Run(":80")
}
