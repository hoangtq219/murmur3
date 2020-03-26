package test

import (
	"fmt"
	"testing"
)

func TestZeroShiftRight(t *testing.T)  {
	var result int64 = 0
	k := 33
	var x int64 = -3984010896604591268
	const n = 64
	if x >= 0 {
		result =  x>>k
	} else {
		result = ( x >> k) + (int64(2) << (63-k))
	}
	fmt.Println(result)
}

func TestLong6(t *testing.T)  {
	val := 8562518959811306728
	//result := make([]byte, 5)
	//result[0] = -44
	output :=  int8(val >> 48)

	fmt.Println(output)
	//if result[0] == ouput {
	//	fmt.Println("true")
	//} else {
	//	fmt.Println("Error")
	//}
}
