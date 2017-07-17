package main

import (
	"JPRO/modules"
	"JPRO/zipping"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	_ "runtime/cgo"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

// how to build the application:
// DEVELOPMENT: main-res.syso main.rc && go build -ldflags="-H windowsgui" -i
// PRODUCTION: main-res.syso main.rc && go build -i

func main() {
	UI()
}

func OnClose() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(sigs, os.Interrupt, os.Kill)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	WriteAll()
	fmt.Println("exiting")
	os.Exit(3)
}

func WriteAll() {
	messages := make(chan int)
	go func() {
		time.Sleep(time.Second * 1)
		zipping.Zipit()
		messages <- 1
	}()
	go func() {
		time.Sleep(time.Second * 1)
		modules.WriteRep(modules.StartTime, modules.EndTime, "log")
		if len(modules.AllTempMax) > 0 {
			modules.WriteMax(modules.AllTempMax)
		}
		messages <- 2
	}()
	go func() {
		for i := range messages {
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Second * 2)
}

func ConfigRuntime() {
	go OnClose()
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	var t *testing.T
	fmt.Println("max: ", modules.MAX)
	go modules.Counter(t)
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Run(":80")
}
