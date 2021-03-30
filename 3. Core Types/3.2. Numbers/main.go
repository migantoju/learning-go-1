package main

import (
	"fmt"
	"runtime"
)

func checkMemStats() {
	// var list []int
	var list []uint8

	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("TotalAlloc (Heap) %v MiB\n", m.TotalAlloc/1024/1024)
}

func main() {
	checkMemStats()
}
