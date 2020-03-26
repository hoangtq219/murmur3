package go_murmur3

import "errors"

type ByteBufferInt8 struct {
	HB              []int8 // data của buffer
	Offset          int
	Mark            int // đánh dấu
	Position        int // vị trí đang lưu byte
	Limit           int // giới hạn
	Capacity        int // khả năng chứa của buffer
	BigEndian       bool
	NativeByteOrder bool
}

// Bits line 395 - 402
func (buf *ByteBufferInt8) putLong(val int64) {
	if buf.BigEndian {
		buf.putLongB(buf.ix(buf.nextPutIndex(8)), val)
	} else {
		buf.putLongL(buf.ix(buf.nextPutIndex(8)), val)
	}
}

// Bits line 373 - 382
func (buf *ByteBufferInt8) putLongB(index int, val int64) {
	buf.HB[index] = Long7Int8(val)
	buf.HB[index+1] = Long6Int8(val)
	buf.HB[index+2] = Long5Int8(val)
	buf.HB[index+3] = Long4Int8(val)
	buf.HB[index+4] = Long3Int8(val)
	buf.HB[index+5] = Long2Int8(val)
	buf.HB[index+6] = Long1Int8(val)
	buf.HB[index+7] = Long0Int8(val)
}

// Bits line 351 - 360
func (buf *ByteBufferInt8) putLongL(index int, val int64) {
	buf.HB[index+7] = Long7Int8(val)
	buf.HB[index+6] = Long6Int8(val)
	buf.HB[index+5] = Long5Int8(val)
	buf.HB[index+4] = Long4Int8(val)
	buf.HB[index+3] = Long3Int8(val)
	buf.HB[index+2] = Long2Int8(val)
	buf.HB[index+1] = Long1Int8(val)
	buf.HB[index] = Long0Int8(val)
}

// Class Buffer line 155 - 163
func (buf *ByteBufferInt8) nextPutIndex(var1 int) int {
	if buf.Limit-buf.Position < var1 {
		errors.New("Exception: Tràn bộ nhớ đệm!")
		return -1
	} else {
		var2 := buf.Position
		buf.Position += var1
		return var2
	}
}

// HeapByteBuffer line 33 - 35
func (buf *ByteBufferInt8) ix(nextPutIndex int) int {
	return buf.Offset + nextPutIndex
}

// HashCodes line 50 - 52
func (buf *ByteBufferInt8) AsInt() int32 {
	// return this.bytes[0] & 255 | (this.bytes[1] & 255) << 8 | (this.bytes[2] & 255) << 16 | (this.bytes[3] & 255) << 24;
	return  int32(buf.HB[0]) & 255 | (int32(buf.HB[1]) & 255) << 8 | (int32(buf.HB[2]) & 255) << 16 | (int32(buf.HB[3]) & 255) << 24
}