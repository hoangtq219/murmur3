package murmur3

import (
	"errors"
)

// Dùng chung để lưu data, vị trí đang putChar (Chung)
type ByteBuffer struct {
	HB              []byte // data của buffer
	Offset          int
	Mark            int // đánh dấu
	Position        int // vị trí đang lưu byte
	Limit           int // giới hạn
	Capacity        int // khả năng chứa của buffer
	BigEndian       bool
	NativeByteOrder bool
}

// Dùng cho khởi tạo buffer khi hash với murmur, theo AbstractStreamingHashFunction line 47
type MurmurByteBuffer struct {
	buffer     *ByteBuffer
	bufferSize int
	chunkSize  int
}

// theo ByteBuffer line 31 - 37
func allocateByteBuffer(bufferSize int) *ByteBuffer {
	hb := make([]byte, bufferSize)
	return &ByteBuffer{
		HB:              hb,
		Offset:          0,
		Mark:            -1,
		Position:        0,
		Limit:           bufferSize,
		Capacity:        bufferSize,
		NativeByteOrder: true,
	}
}

// theo AbstractStreamingHasher, line 52 - 60
func InitMurmurByteBuffer(chunkSize int) *MurmurByteBuffer {
	mBuff := &MurmurByteBuffer{}
	mBuff.chunkSize = chunkSize // AbstractStreamingHasher line 58 - 59
	mBuff.bufferSize = chunkSize
	// 16 + 7 = 23 AbstractStreamingHasher line 57
	mBuff.buffer = allocateByteBuffer(mBuff.bufferSize + 7)
	return mBuff
}

// 
func (buf *ByteBuffer) putCharL(index int, val byte) {
	buf.HB[index] = val
	buf.HB[index+1] = val >> 8
	//buf.Position += 2
}

// theo AbstractStreamingHashFunction line 163
func (buf *ByteBuffer) remaining() int {
	return buf.Limit - buf.Position
}

func (buf *ByteBuffer) flip() {
	buf.Limit = buf.Position
	buf.Position = 0
	buf.Mark = -1
}

// Bits line 303 - 305
func (buf *ByteBuffer) getLongB(index int) int64 {
	var0 := getByte(buf.HB, index)
	var1 := getByte(buf.HB, index+1)
	var2 := getByte(buf.HB, index+2)
	var3 := getByte(buf.HB, index+3)
	var4 := getByte(buf.HB, index+4)
	var5 := getByte(buf.HB, index+5)
	var6 := getByte(buf.HB, index+6)
	var7 := getByte(buf.HB, index+7)

	return makeLong(var0, var1, var2, var3, var4, var5, var6, var7)
}

// Bits line 295 - 297
func (buf *ByteBuffer) getLongL(index int) int64 {
	var0 := getByte(buf.HB, index+7)
	var1 := getByte(buf.HB, index+6)
	var2 := getByte(buf.HB, index+5)
	var3 := getByte(buf.HB, index+4)
	var4 := getByte(buf.HB, index+3)
	var5 := getByte(buf.HB, index+2)
	var6 := getByte(buf.HB, index+1)
	var7 := getByte(buf.HB, index)

	return makeLong(var0, var1, var2, var3, var4, var5, var6, var7)
}

// Buffer line 137 - 145
func (buf *ByteBuffer) nextGetIndex(var1 int) int {
	if buf.Limit-buf.Position < var1 {
		errors.New("Buffer Underflow Exception!")
	} else {
		var2 := buf.Position
		buf.Position += var1
		return var2
	}
	return -1
}

// HeapBuffer 203 - 205
func (buf *ByteBuffer) getLong() int64 {
	if buf.BigEndian {
		return buf.getLongB(buf.ix(buf.nextGetIndex(8)))
	} else {
		return buf.getLongL(buf.ix(buf.nextGetIndex(8)))
	}
}

func (buf *ByteBuffer) capacity() int {
	return buf.Capacity
}

func (buf *ByteBuffer) discardMark() {
	buf.Mark = -1
}

func (buf *ByteBuffer) compact() {
	buf.position(buf.remaining())
	buf.limit(buf.capacity())
	buf.discardMark()
}

func (buf *ByteBuffer) get(index int) byte {
	return buf.HB[index]
}

func getByte(buffer []byte, index int) byte {
	return buffer[index]
}

func makeLong(var0, var1, var2, var3, var4, var5, var6, var7 byte) int64 {
	return int64(var0)<<56 | (int64(var1)&255)<<48 | (int64(var2)&255)<<40 | (int64(var3)&255)<<32 | (int64(var4)&255)<<24 | (int64(var5)&255)<<16 | (int64(var6)&255)<<8 | (int64(var7))&255
}

// Class Buffer 42 - 53
func (buf *ByteBuffer) position(var1 int) {
	if var1 <= buf.Limit && var1 >= 0 {
		buf.Position = var1
		if buf.Mark > buf.Position {
			buf.Mark = -1
		}
	} else {
		errors.New("Exception: Đối số truyền vào không hợp lệ!!!")
	}
}

// Buffer line 59 - 74
func (buf *ByteBuffer) limit(chunkSize int) {
	if chunkSize <= buf.Capacity && chunkSize >= 0 {
		buf.Limit = chunkSize
		if buf.Position > buf.Limit {
			buf.Position = buf.Limit
		}

		if buf.Mark > buf.Limit {
			buf.Mark = -1
		}
	} else {
		errors.New("Exception: Đối số truyền vào không hợp lệ!")
	}
}

// HeapByteBuffer line 33 - 35
func (buf *ByteBuffer) ix(nextPutIndex int) int {
	return buf.Offset + nextPutIndex
}

// Class Buffer line 155 - 163
func (buf *ByteBuffer) nextPutIndex(var1 int) int {
	if buf.Limit-buf.Position < var1 {
		errors.New("Exception: Tràn bộ nhớ đệm!")
		return -1
	} else {
		var2 := buf.Position
		buf.Position += var1
		return var2
	}
}

// Bits line 373 - 382
func (buf *ByteBuffer) putLongB(index int, val int64) {
	buf.HB[index] = Long7(val)
	buf.HB[index+1] = Long6(val)
	buf.HB[index+2] = Long5(val)
	buf.HB[index+3] = Long4(val)
	buf.HB[index+4] = Long3(val)
	buf.HB[index+5] = Long2(val)
	buf.HB[index+6] = Long1(val)
	buf.HB[index+7] = Long0(val)
}

// Bits line 351 - 360
func (buf *ByteBuffer) putLongL(index int, val int64) {
	buf.HB[index+7] = Long7(val)
	buf.HB[index+6] = Long6(val)
	buf.HB[index+5] = Long5(val)
	buf.HB[index+4] = Long4(val)
	buf.HB[index+3] = Long3(val)
	buf.HB[index+2] = Long2(val)
	buf.HB[index+1] = Long1(val)
	buf.HB[index] = Long0(val)
}

// Bits line 395 - 402
func (buf *ByteBuffer) putLong(val int64) {
	if buf.BigEndian {
		buf.putLongB(buf.ix(buf.nextPutIndex(8)), val)
	} else {
		buf.putLongL(buf.ix(buf.nextPutIndex(8)), val)
	}
}

// Buffer line 165 - 171
func (buf *ByteBuffer) checkIndex(index int) int {
	if index >= 0 && index < buf.Limit {
		return index
	} else {
		errors.New("Exception: Vượt index trong giới hạn mảng")
	}
	return -1
}

// Class Buffer 38 - 53
func (mBB *MurmurByteBuffer) positionFunc(limit int) {
	mBB.buffer.position(limit)
}

func (mBB *MurmurByteBuffer) limitFunc(chunkSize int) {
	mBB.buffer.limit(chunkSize)
}

func (mBB *MurmurByteBuffer) limit() int {
	return mBB.buffer.Limit
}

func (mBB *MurmurByteBuffer) position() int {
	return mBB.buffer.Position
}

func (mBB *MurmurByteBuffer) remaining() int {
	return mBB.buffer.Limit - mBB.buffer.Position
}

// Bits line 395 - 402
func (mBB *MurmurByteBuffer) putLong(val int64) {
	mBB.buffer.putLong(val)
}

// Buffer line 98 - 103
func (mBB *MurmurByteBuffer) flip() {
	mBB.buffer.flip()
}

func (mBB *MurmurByteBuffer) get(index int) byte {
	return mBB.buffer.get(mBB.buffer.ix(mBB.buffer.checkIndex(index)))
}

func (mBB *MurmurByteBuffer) getLong() int64 {
	return mBB.buffer.getLong()
}

func (mBB *MurmurByteBuffer) putLongMur3(val int64) {
	if mBB.buffer.BigEndian {
		mBB.buffer.putLongB(mBB.buffer.ix(mBB.buffer.nextPutIndex(8)), val)
	} else {
		mBB.buffer.putLongL(mBB.buffer.ix(mBB.buffer.nextPutIndex(8)), val)
	}
}

// HashCodes line 50 - 52
func (buf *ByteBuffer) AsInt() int32 {
	// return this.bytes[0] & 255 | (this.bytes[1] & 255) << 8 | (this.bytes[2] & 255) << 16 | (this.bytes[3] & 255) << 24;
	return  int32(buf.HB[0]) & 255 | (int32(buf.HB[1]) & 255) << 8 | (int32(buf.HB[2]) & 255) << 16 | (int32(buf.HB[3]) & 255) << 24
}

func (buf *ByteBuffer) ToBytes(val int) []byte {
	b := make([]byte, 4)

	for i := 3; i > 0; i-- {
		b[i] = byte(val)
		val = rightShift(val, 8)
	}

	b[0] = byte(val)
	return b
}

func rightShift(x int, k int) int {
	const n = 64
	if x >= 0 {
		return x>>k
	} else {
		return ( x >> k) + (int(2) << (63-k))
	}
}