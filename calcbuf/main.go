package main

import "fmt"

func main() {
	// TODO serve if -s/p command
	serve()
	}


func example(){
 	cb := CalcBuf{OneRepMax: 195, Buffer: 2} // my last DL PR
	result := cb.RMBufferSeries(cb.OneRepMax, cb.Buffer, nil)

	fmt.Println(result)
}
