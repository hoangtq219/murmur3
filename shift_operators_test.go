package murmur3

import (
	"fmt"
	"testing"
)

func TestZeroShiftRight(t *testing.T) {
	var result int64 = 0
	k := uint(33)
	var x int64 = -3984010896604591268
	if x >= 0 {
		result = x >> k
	} else {
		result = (x >> k) + (int64(2) << (63 - k))
	}
	fmt.Println(result)
}

func TestLong6(t *testing.T) {
	val := 8562518959811306728

	output := int8(val >> 48)

	fmt.Println(output)
	if output == -44 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func TestIntToBytes(t *testing.T) {
	postId := "3546546464654654"
	fmt.Println(IntToBytes(postId))
}