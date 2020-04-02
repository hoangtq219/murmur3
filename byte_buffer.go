package murmur3

type ByteBuffer struct {
	HB              []byte
	Offset          int
	Mark            int
	Position        int
	Limit           int
	Capacity        int
	BigEndian       bool
	NativeByteOrder bool
}

func allocateByteBuffer(bufferSize int) *ByteBuffer {
	return &ByteBuffer{
		HB:              make([]byte, bufferSize),
		Offset:          0,
		Mark:            -1,
		Position:        0,
		Limit:           bufferSize,
		Capacity:        bufferSize,
		NativeByteOrder: true,
	}
}

func (buf *ByteBuffer) putCharL(index int, val byte) {
	buf.HB[index] = val
	buf.HB[index+1] = val >> 8
}

func (buf *ByteBuffer) remaining() int {
	return buf.Limit - buf.Position
}

func (buf *ByteBuffer) flip() {
	buf.Limit = buf.Position
	buf.Position = 0
	buf.Mark = -1
}

func (buf *ByteBuffer) getLongB(index int) int64 {
	return makeLong(get(buf.HB, index), get(buf.HB, index+1), get(buf.HB, index+2), get(buf.HB, index+3), get(buf.HB, index+4), get(buf.HB, index+5), get(buf.HB, index+6), get(buf.HB, index+7))
}

func (buf *ByteBuffer) getLongL(index int) int64 {
	return makeLong(get(buf.HB, index+7), get(buf.HB, index+6), get(buf.HB, index+5), get(buf.HB, index+4), get(buf.HB, index+3), get(buf.HB, index+2), get(buf.HB, index+1), get(buf.HB, index))
}

func (buf *ByteBuffer) nextGetIndex(var1 int) int {
	if buf.Limit-buf.Position < var1 {
		handleWarnPrintf("Buffer Underflow Exception")
	} else {
		var2 := buf.Position
		buf.Position += var1
		return var2
	}
	return -1
}

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

func (buf *ByteBuffer) position(var1 int) {
	if var1 <= buf.Limit && var1 >= 0 {
		buf.Position = var1
		if buf.Mark > buf.Position {
			buf.Mark = -1
		}
	} else {
		handleWarnPrintf("Illegal Argument Exception")
	}
}

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
		handleWarnPrintf("Illegal Argument Exception")
	}
}

func (buf *ByteBuffer) ix(nextPutIndex int) int {
	return buf.Offset + nextPutIndex
}

func (buf *ByteBuffer) nextPutIndex(var1 int) int {
	if buf.Limit-buf.Position < var1 {
		handleWarnPrintf("Buffer Overflow Exception!")
		return -1
	} else {
		var2 := buf.Position
		buf.Position += var1
		return var2
	}
}

func (buf *ByteBuffer) putLongB(index int, val int64) {
	buf.HB[index] = long7(val)
	buf.HB[index+1] = long6(val)
	buf.HB[index+2] = long5(val)
	buf.HB[index+3] = long4(val)
	buf.HB[index+4] = long3(val)
	buf.HB[index+5] = long2(val)
	buf.HB[index+6] = long1(val)
	buf.HB[index+7] = long0(val)
}

func (buf *ByteBuffer) putLongL(index int, val int64) {
	buf.HB[index+7] = long7(val)
	buf.HB[index+6] = long6(val)
	buf.HB[index+5] = long5(val)
	buf.HB[index+4] = long4(val)
	buf.HB[index+3] = long3(val)
	buf.HB[index+2] = long2(val)
	buf.HB[index+1] = long1(val)
	buf.HB[index] = long0(val)
}

// Puts a long into this sink.
func (buf *ByteBuffer) putLong(val int64) {
	if buf.BigEndian {
		buf.putLongB(buf.ix(buf.nextPutIndex(8)), val)
	} else {
		buf.putLongL(buf.ix(buf.nextPutIndex(8)), val)
	}
}

func (buf *ByteBuffer) checkIndex(index int) int {
	if index >= 0 && index < buf.Limit {
		return index
	} else {
		handleWarnPrintf("Array Index Out Of Bounds Exception")
	}
	return -1
}

// AsInt() Returns the first four bytes of this hashcode's byte, converted to an int value little-endian order
func (buf *ByteBuffer) AsInt() int {
	return int(int32(buf.HB[0])&255 | (int32(buf.HB[1])&255)<<8 | (int32(buf.HB[2])&255)<<16 | (int32(buf.HB[3])&255)<<24)
}

// AsBytes() Returns the first four bytes of this hash code as a byte array.
func (buf *ByteBuffer) AsBytes() []byte {
	return []byte{buf.HB[3], buf.HB[2], buf.HB[1], buf.HB[0]}
}

// AsBytes() Returns the value of this hash code as a byte array.
func (buf *ByteBuffer) ToBytes() []byte {
	return buf.HB
}

// ToString() Returns a string containing each byte of ToBytes(), in order, as a two-digit unsigned hexadecimal number in lower case.
// Note that if the output is considered to be a single hexadecimal number, this hash code's bytes are the big-endian representation of that number.
func (buf *ByteBuffer) ToString() string {
	return toString(buf.HB)
}
