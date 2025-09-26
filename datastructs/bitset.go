package datastructs

const bitsInBlock = 64

type bitSet struct {
	bits []uint64
}

func NewBitSet(size uint64) *bitSet {
	return &bitSet{
		bits: make([]uint64, size/bitsInBlock),
	}
}

func (b *bitSet) Add(value uint32) {
	idx := value / bitsInBlock
	pos := value % bitsInBlock
	b.bits[idx] |= 1 << pos
}

func (b *bitSet) IsBitExist(value uint32) bool {
	idx := value / bitsInBlock
	pos := value % bitsInBlock
	return b.bits[idx]&(1<<pos) != 0
}
