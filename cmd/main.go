package main

import (
	"log"
	"murmur3"
)

func main()  {
	m3_128 := murmur3.HashString(-1467523828, "681236075540516864")

	log.Println(m3_128.AsBytes())
	log.Println(m3_128.ToBytes())
	log.Println(m3_128.AsInt())
	log.Println(m3_128.ToString())

	//key := append(m3_128.AsBytes(), murmur3.IntToBytes("681236075540516864")...)
	//key = append(key, murmur3.IntToBytes("17480678941457235")...)
	//fmt.Println(key)
	//key1 := []byte{182, 25, 164, 232, 9, 116, 60, 153, 242, 198, 16, 0, 0, 62, 26, 149, 186, 188, 251, 83}
	//fmt.Println(key1)

	//fmt.Println(m3_128.AsInt())
	//fmt.Println(m3_128.ToString())
}
