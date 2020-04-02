package main

import (
	"fmt"
	"log"
	"murmur3"
)

func main() {

	m3_128 := murmur3.HashString(-121254478, "681236075540516864")

	log.Println(m3_128.AsInt())    // 257976075
	log.Println(m3_128.AsBytes())  // [15 96 103 11]
	log.Println(m3_128.ToBytes())  // [11 103 96 15 2 251 20 14 21 75 88 130 194 211 36 22]
	log.Println(m3_128.ToString()) //  0b67600f02fb140e154b5882c2d32416

	key := append(m3_128.AsBytes(), murmur3.IntToBytes("681236075540516864")...)
	key = append(key, murmur3.IntToBytes("17480678941457235")...)
	fmt.Println(key) // [15 96 103 11 9 116 60 153 242 198 16 0 0 62 26 149 186 188 251 83]
}
