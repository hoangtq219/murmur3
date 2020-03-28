package murmur3

import (
	"log"
	"runtime"
	"strconv"
	"strings"
)

func HandlePrintf(msg interface{}) {
	log.Printf("[I] %v", msg)
}

func HandleWarnPrintf(msg interface{}) {
	log.Printf("[W] %v", msg)
}

func HandleErrorFatalf(err interface{}) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Fatalf("[E] %v %s:%d", err, fn, line)
	}
}

func ToBytes(val int64) []byte {
	b := make([]byte, 8)
	for i := 7; i > 0; i-- {
		b[i] = byte(val)
		val = RightShift(val, 8)
	}
	b[0] = byte(val)
	return b
}

/*
	Chuyển 16 byte sau khi hash thành chuỗi dài 32 byte
*/
func ToString(bytes []byte) string {
	var output strings.Builder
	for i := 0; i < len(bytes); i++ {
		b := int8(bytes[i])
		output.WriteByte(HexDigits[b >> 4 & 15])
		output.WriteByte(HexDigits[b & 15])
	}
	return output.String()
}

func RotateLeft(x int64, k int) int64 {
	const n = 64
	return x<<k | int64(uint64(x)>>(n-k))
}

/*
n >>> s
	* if n is positive, then the result is the same as that of n >> s
	* if n is negative and the type of the left-hand operand is int, then the result is equal to that of the expression (n >> s) + (2 << ~s)
	* if n is negative and the type of the left-hand operand is long, then the result is equal to that of the expression (n >> s) + (int64(2) << ~s)

Bonus:
	* ~s as a shift distance is equivalent to `31 - s` when shifting an int value and to `63 - s` when shifting a long value
*/
func RightShift(x int64, k int) int64 {
	if x >= 0 {
		return x>>k
	} else {
		return ( x >> k) + (int64(2) << (63-k))
	}
}

func Long0(val int64) byte {
	return byte(int(val))
}

func Long1(val int64) byte {
	return byte(int(val >> 8))
}

func Long2(val int64) byte {
	return byte(int(val >> 16))
}

func Long3(val int64) byte {
	return byte(int(val >> 24))
}

func Long4(val int64) byte {
	return byte(int(val >> 32))
}

func Long5(val int64) byte {
	return byte(int(val >> 40))
}

func Long6(val int64) byte {
	return byte(int(val >> 48))
}

func Long7(val int64) byte {
	return byte(int(val >> 56))
}


func Get(buffer []byte, index int) byte {
	return buffer[index]
}

func MakeLong(var0, var1, var2, var3, var4, var5, var6, var7 byte) int64 {
	return int64(var0)<<56 | (int64(var1)&255)<<48 | (int64(var2)&255)<<40 | (int64(var3)&255)<<32 | (int64(var4)&255)<<24 | (int64(var5)&255)<<16 | (int64(var6)&255)<<8 | (int64(var7))&255
}

func IntToBytes(input string) []byte {
	number, _ := strconv.Atoi(input)
	return ToBytes(int64(number))
}