package main

import (
	"fmt"
	"murmur3"
)

func main()  {
	m3_128 := murmur3.HashString(-1467523828, "681236075540516864")

	fmt.Println(m3_128.AsInt())
	fmt.Println(m3_128.ToString())
}
