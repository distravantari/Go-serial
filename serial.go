package main

import (
	"strings"
	"testing"
	"time"

	"github.com/tarm/serial"

	"fmt"
)

func Counter(t *testing.T) {
	port0 := "COM2"
	port1 := "COM1"
	if port0 == "" || port1 == "" {
		t.Skip("Skipping test because PORT0 or PORT1 environment variable is not set")
	}
	c0 := &serial.Config{Name: port0, Baud: 2400}
	c1 := &serial.Config{Name: port1, Baud: 2400}

	s1, err := serial.OpenPort(c0)
	if err != nil {
		fmt.Println(err)
	}

	s2, err := serial.OpenPort(c1)
	if err != nil {
		fmt.Println(err)
	}

	ch := make(chan int, 1)
	go func() {
		buf := make([]byte, 128)
		var readCount int
		for {
			n, err := s2.Read(buf)
			if err != nil {
				fmt.Print(err)
			}
			readCount++
			// fmt.Printf("Read %v %v bytes: % 02x %s", readCount, n, buf[:n], buf[:n])
			weight := fmt.Sprintf("%s ", buf[:n])
			if strings.ContainsAny(weight, "+") {
				weight = strings.Trim(weight, "+")
				fmt.Print(weight + " ")
			}

			select {
			case <-ch:
				ch <- readCount
				close(ch)
			default:
			}
		}
	}()

	if _, err = s1.Write([]byte(" ")); err != nil {
		fmt.Println(err)
	}
	if _, err = s1.Write([]byte(" ")); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond / 1)
	if _, err = s1.Write([]byte(" ")); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond / 1)

	ch <- 0
	s1.Write([]byte(" ")) // We could be blocked in the read without this
	c := <-ch
	exp := 5
	if c >= exp {
		fmt.Println("Expected less than %v read, got %v", exp, c)
	}
}

func byteSlice(arr []byte) byte {
	sum := byte(0)
	for _, b := range arr {
		sum += b
	}
	return sum
}
