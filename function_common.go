package murmur3

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

func ToBytes(val int) []byte {
	b := make([]byte, 8)

	for i := 7; i > 0; i-- {
		b[i] = byte(val)
		val = RightShift(val, 8)
	}

	b[0] = byte(val)
	return b
}

func RightShift(x int, k int) int {
	if x >= 0 {
		return x>>k
	} else {
		return ( x >> k) + (int(2) << (63-k))
	}
}