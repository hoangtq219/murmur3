package murmur3

import (
	"fmt"
	"testing"
)

func TestHashString(t *testing.T) {
	m3_128 := HashString(-121254478, "681236075540516864")
	fmt.Println(m3_128.AsInt())
	fmt.Println(m3_128.AsIntBytes())
	fmt.Println(m3_128.ToBytes())
	fmt.Println(m3_128.ToString())
}
