Go Murmur3 128 
=== 

> The project is building to solve the hash with seed (positive and negative) in golang.  
> You can get the code for implementing it on your project.  

Output
---

* `Integer:` 
  * `AsInt()` Returns the first four bytes of this hashcode's byte, converted to an int value little-endian order  
* `String:` 
  * `ToString()` Returns a string containing each byte of `ToBytes()`, in order, as a two-digit unsigned hexadecimal number in lower case.
  * Note that if the output is considered to be a single hexadecimal number, this hash code's bytes are the big-endian representation of that number. 
* `Bytes:` 
  * `ToBytes()` Returns the value of this hash code as a byte array.
  * `AsIntBytes()` Returns the first four bytes of this hash code as a byte array.
  * `AsLongBytes()` Returns the first eight bytes of this hash code as a byte array.
  
Examples: 

```go
m3_128 := murmur3.HashString(-121254478, "681236075540516864")
 
// AsInt()
m3_128.AsInt() 
// output: 257976075

// AsBytes() 
m3_128.AsIntBytes() 
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

key := append(m3_128.AsIntBytes(), murmur3.IntToBytes("681236075540516864")...)
key = append(key, murmur3.IntToBytes("17480678941457235")...)
fmt.Println(key)
// output: [15 96 103 11 9 116 60 153 242 198 16 0 0 62 26 149 186 188 251 83]
```

Benchmarks
---

> Core i5 5200, 2.4 Ghz. All runs include hasher instantiation and sequence finalization.

```text
BenchmarkHashString128/1-4               8497768               284 ns/op           3.53 MB/s         128 B/op          3 allocs/op
BenchmarkHashString128/2-4               7512295               320 ns/op           6.25 MB/s         128 B/op          3 allocs/op
BenchmarkHashString128/4-4               7209426               330 ns/op          12.12 MB/s         128 B/op          3 allocs/op
BenchmarkHashString128/8-4               5748768               414 ns/op          19.34 MB/s         128 B/op          3 allocs/op
BenchmarkHashString128/16-4              4025712               594 ns/op          26.92 MB/s         128 B/op          3 allocs/op
BenchmarkHashString128/32-4              2567530               937 ns/op          34.15 MB/s         128 B/op          3 allocs/op
BenchmarkHashString128/64-4              1429455              1680 ns/op          38.10 MB/s         192 B/op          4 allocs/op
BenchmarkHashString128/128-4              726985              3056 ns/op          41.88 MB/s         256 B/op          4 allocs/op
BenchmarkHashString128/256-4              402700              5844 ns/op          43.80 MB/s         384 B/op          4 allocs/op
BenchmarkHashString128/512-4              205081             11674 ns/op          43.86 MB/s         640 B/op          4 allocs/op
BenchmarkHashString128/1024-4             102885             22581 ns/op          45.35 MB/s        1152 B/op          4 allocs/op
BenchmarkHashString128/2048-4              53106             44812 ns/op          45.70 MB/s        2176 B/op          4 allocs/op
BenchmarkHashString128/4096-4              26314             90024 ns/op          45.50 MB/s        4224 B/op          4 allocs/op
BenchmarkHashString128/8192-4              13608            179162 ns/op          45.72 MB/s        8320 B/op          4 allocs/op
     
```

Compatibility
---

> Minimum go version 1.10

Contributing to murmur3
---

> Just make pull request. You are in!

LICENSE 
---

Copyright (c) 2020 Platform. See the `LICENSE` file for more info.