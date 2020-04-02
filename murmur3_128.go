package murmur3

import (
	"log"
)

/*
   ###############################
   ##     guava version13.0     ##
   ###############################
*/

type Murmur3_128Hasher struct {
	H1     int64
	H2     int64
	Length int
	bb     *MurmurByteBuffer
}

func initMurmur3_128Hsher(seed int64) *Murmur3_128Hasher {
	return &Murmur3_128Hasher{
		H1:     seed,
		H2:     seed,
		Length: 0,
		bb:     initMurmurByteBuffer(CHUNK_SIZE),
	}
}

func (m *Murmur3_128Hasher) putString(CharSequence string) {
	for i := 0; i < len(CharSequence); i++ {
		//log.Printf("position: %d, limit: %d", m.bb.buffer.position, m.bb.buffer.limit)
		m.putChar(CharSequence[i])
		m.munchIfFull()
	}
}

// Puts a character into this sink.
func (m *Murmur3_128Hasher) putChar(val byte) {
	//HandlePrintf("putChar: " + strconv.FormatInt(int64(index), 10))
	m.bb.putCharL(m.bb.ix(m.bb.nextPutIndex(2)), val)
}

func (m *Murmur3_128Hasher) munchIfFull() {
	if m.bb.remaining() < 8 {
		m.munch()
	}
}

func (m *Murmur3_128Hasher) munch() {
	//log.Printf("remaining = %d, position: %d, limit: %d", m.bb.remaining(), m.bb.position(), m.bb.limit())
	m.bb.flip()
	for m.bb.remaining() >= m.bb.chunkSize {
		m.process()
	}
	m.bb.buffer.compact()
}

func (m *Murmur3_128Hasher) process() {
	k1 := m.bb.getLong()
	k2 := m.bb.getLong()

	m.bmix64(k1, k2)
	m.Length += 16
}

func (m *Murmur3_128Hasher) bmix64(k1, k2 int64) {
	m.H1 ^= mixK1(k1)
	m.H1 = rotateLeft(m.H1, 27)
	m.H1 += m.H2
	m.H1 = m.H1*5 + 1390208809
	m.H2 ^= mixK2(k2)
	m.H2 = rotateLeft(m.H2, 31)
	m.H2 += m.H1
	m.H2 = m.H2*5 + 944331445
}

func (m *Murmur3_128Hasher) processRemainingAfterBmixData() {
	m.bb.positionFunc(m.bb.limit())
	m.bb.limitFunc(m.bb.chunkSize + 7)

	for m.bb.position() < m.bb.chunkSize {
		m.bb.putLong(0)
	}
	m.bb.limitFunc(m.bb.chunkSize)
	m.bb.flip()
	m.process()
}

// HashString returns a 128 bits hasher set with explicit seed value
func HashString(seed int64, data string) *ByteBuffer {
	m3 := initMurmur3_128Hsher(seed)
	m3.putString(data)

	m3.munch()
	m3.bb.buffer.flip()
	if m3.bb.remaining() > 0 {
		m3.processRemaining()
	}
	return m3.makeHash()
}

// Computes a hash code based on the data that have been provided to this hasher
func (m *Murmur3_128Hasher) makeHash() *ByteBuffer {
	m.H1 ^= int64(m.Length)
	m.H2 ^= int64(m.Length)
	m.H1 += m.H2
	m.H2 += m.H1
	m.H1 = fmix64(m.H1)
	m.H2 = fmix64(m.H2)
	m.H1 += m.H2
	m.H2 += m.H1

	bb := &ByteBuffer{
		// isReadOnly false
		HB:              make([]byte, 16),
		Offset:          0,
		Mark:            -1,
		Position:        0,
		Limit:           16,
		Capacity:        16,
		BigEndian:       false,
		NativeByteOrder: false,
	}

	bb.putLong(m.H1)
	bb.putLong(m.H2)
	return bb
}

// Finalization mix - force all bits of a hash block to avalanche
func fmix64(k int64) int64 {
	k = k ^ rightShift(k, 33)
	k *= -49064778989728563
	k = k ^ rightShift(k, 33)
	k *= -4265267296055464877
	k = k ^ rightShift(k, 33)
	return k
}

func mixK1(k1 int64) int64 {
	k1 *= c1
	k1 = rotateLeft(k1, 31)
	k1 *= c2
	return k1
}

func mixK2(k2 int64) int64 {
	k2 *= c2
	k2 = rotateLeft(k2, 33)
	k2 *= c1
	return k2
}

/*
	Sau khi kết thúc quá trình putChar, nếu số bytes còn lại > 0 (các byte này vẫn chưa được process với bmix64)
tiến hành process remaining với số bytes còn lại này trước khi hash
*/
func (m *Murmur3_128Hasher) processRemaining() {
	var k1 int64 = 0
	var k2 int64 = 0

	m.Length += m.bb.remaining()

	switch m.bb.remaining() {
	case 7:
		k1 ^= int64(m.bb.get(6)) << 48
		fallthrough
	case 6:
		k1 ^= int64(m.bb.get(5)) << 40
		fallthrough
	case 5:
		k1 ^= int64(m.bb.get(4)) << 32
		fallthrough
	case 4:
		k1 ^= int64(m.bb.get(3)) << 24
		fallthrough
	case 3:
		k1 ^= int64(m.bb.get(2)) << 16
		fallthrough
	case 2:
		k1 ^= int64(m.bb.get(1)) << 8
		fallthrough
	case 1:
		k1 ^= int64(m.bb.get(0))
		break
	case 15:
		k2 ^= int64(m.bb.get(14)) << 48
		fallthrough
	case 14:
		k2 ^= int64(m.bb.get(13)) << 40
		fallthrough
	case 13:
		k2 ^= int64(m.bb.get(12)) << 32
		fallthrough
	case 12:
		k2 ^= int64(m.bb.get(11)) << 24
		fallthrough
	case 11:
		k2 ^= int64(m.bb.get(10)) << 16
		fallthrough
	case 10:
		k2 ^= int64(m.bb.get(9)) << 8
		fallthrough
	case 9:
		k2 ^= int64(m.bb.get(8))
		m.H2 ^= mixK2(k2)
		fallthrough
	case 8:
		k1 ^= m.bb.getLong()
		break
	default:
		log.Println("[E] Should never get here")
	}

	m.H1 ^= mixK1(k1)
	m.H2 ^= mixK2(k2)
}
