package main

import (
	"fmt"
	go_murmur3 "go-murmur3"
)

func main()  {
	m3_128 := go_murmur3.HashString(-1467523828, "681236075540516864")
	fmt.Println(m3_128.AsInt())
}
