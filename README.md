Go Murmur3 
=== 

> Project được xây dựng nhằm mục đích giải quyết việc hash với seed là số âm trong golang  

Ouput
---

* `Integer:` 
  * `AsInt()` Trả về 1 số nguyên 32 bits được lấy từ 4 bytes đầu tiên của `ByteBuffer`  
* `String:` 
  * `ToString()` Trả về 1 chuỗi 256 bits, được hash từ 16 bytes của `ByteBuffer`  
* `Bytes:` 
  * `ToBytes()` Trả về mảng 16 bytes từ `H1` và `H2` sau khi hash
  * `AsBytes()` Trả về mảng 4 bytes đầu tiên của ToBytes theo thứ tự ngược lại
  
Examples: 

```go
m3_128 := murmur3.HashString(-121254478, "681236075540516864")
 
// AsInt()
m3_128.AsInt() 
// output: 257976075

// AsBytes() 
m3_128.AsBytes() 
// output:[15 96 103 11]

// ToBytes() 
m3_128.ToBytes() 
// output: [11 103 96 15 2 251 20 14 21 75 88 130 194 211 36 22]

// ToString()
m3_128.ToString() 
// output: 0b67600f02fb140e154b5882c2d32416

```

**Gen key**  

> Tạo key gồm 20 bytes 

```go
m3_128 := murmur3.HashString(-121254478, "681236075540516864")

key := append(m3_128.AsBytes(), murmur3.IntToBytes("681236075540516864")...)
key = append(key, murmur3.IntToBytes("17480678941457235")...)
fmt.Println(key)
// output: [15 96 103 11 9 116 60 153 242 198 16 0 0 62 26 149 186 188 251 83]
```

LICENSE 
---

Copyright (c) 2020 Platform BigData Adtech Admicro