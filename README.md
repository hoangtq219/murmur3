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
m3_128 := murmur3.HashString(-1467523828, "681236075540516864")
 
// AsInt()
m3_128.AsInt() 
// output: -1239833368

// ToString()
m3_128.ToString() 
// output: e8a419b6a02dd4769f55b02614e8644e 

// ToBytes() 
m3_128.ToBytes() 
// output: [232 164 25 182 160 45 212 118 159 85 176 38 20 232 100 78]

// AsBytes() 
m3_128.AsBytes() 
// output: [182 25 164 232]
```

