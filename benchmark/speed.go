package main

import (
	"fmt"
	"time"
	tx "github.com/valeralabs/sdk/transaction"
)

var raw []byte = []byte("0000000001040015c31b8c1c11c515e244b75806bac48d1399c775000000000000000a0000000000000001000135230d6a06749377212758f925871c1106fb258abed738df874a858f6f394881786d413dccfe8b134f0a8f820ce72220b8201b1cfc43ac365cf1e23f4f4003ab030200000000000516f5f1d7c482c6e1f8330b4805c5c03d8409d32452000000000000006468656c6c6f0000000000000000000000000000000000000000000000000000000000")

func main() {
	// benchmarking the speed of unmarshalling a transaction

	var times []int64

	for i := 0; i < 500000; i++ {
		start := time.Now()
		test()
		end := time.Now()
		fmt.Printf("%d: %v\n", i, end.Sub(start))
		times = append(times, end.Sub(start).Nanoseconds())
	}

	// print the average time
	var total int64
	for _, t := range times {
		total += t
	}
	fmt.Printf("average: %v\n", time.Duration(total/int64(len(times))))
}

func test() {
	var transaction tx.StacksTransaction
	transaction.Unmarshal(raw)
}