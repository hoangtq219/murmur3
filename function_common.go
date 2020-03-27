package murmur3

import "strings"

//func HandlePrintf(msg interface{}) {
//	log.Printf("[I] %v", msg)
//}
//
//func HandleWarnPrintf(msg interface{}) {
//	log.Printf("[W] %v", msg)
//}
//
//func HandleErrorFatalf(err interface{}) {
//	if err != nil {
//		_, fn, line, _ := runtime.Caller(1)
//		log.Fatalf("[E] %v %s:%d", err, fn, line)
//	}
//}

const (
	HexDigits = "0123456789abcdef"
)

func ToBytes(val int64) []byte {
	b := make([]byte, 8)

	for i := 7; i > 0; i-- {
		b[i] = byte(val)
		val = RightShift(val, 8)
	}

	b[0] = byte(val)
	return b
}

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

func RightShift(x int64, k int) int64 {
	if x >= 0 {
		return x>>k
	} else {
		return ( x >> k) + (int64(2) << (63-k))
	}
}