package murmur3

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