package main

import "fmt"


func main() {
 	cb := CalcBuf{OneRepMax: 190, Buffer: 2} // my last DL PR
	result := cb.CalculateBufferSeries(cb.OneRepMax, cb.Buffer, nil)

	fmt.Println(result)
	}
