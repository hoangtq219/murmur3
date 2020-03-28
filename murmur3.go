package murmur3

const (
	CHUNK_SIZE = 16
	c1         = -8663945395140668459
	c2         = 5545529020109919103
	HexDigits  = "0123456789abcdef"
)

type MurmurByteBuffer struct {
	buffer     *ByteBuffer
	bufferSize int
	chunkSize  int
}

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

func (mBB *MurmurByteBuffer) putLong(val int64) {
	mBB.buffer.putLong(val)
}

func (mBB *MurmurByteBuffer) ix(nextPutIndex int) int {
	return mBB.buffer.Offset + nextPutIndex
}

func (mBB *MurmurByteBuffer) nextPutIndex(var1 int) int {
	return mBB.buffer.nextPutIndex(var1)
}

func (mBB *MurmurByteBuffer) putCharL(index int, val byte) {
	mBB.buffer.putCharL(index, val)
}

func (mBB *MurmurByteBuffer) flip() {
	mBB.buffer.flip()
}

func (mBB *MurmurByteBuffer) get(index int) byte {
	return mBB.buffer.get(mBB.buffer.ix(mBB.buffer.checkIndex(index)))
}

func (mBB *MurmurByteBuffer) getLong() int64 {
	return mBB.buffer.getLong()
}
