package bitmap

type BitMap struct{
	bm map[uint64]uint64
}

const bitlen uint64 = 64

func NewBitMap() *BitMap {
	return &BitMap{
		bm: make(map[uint64]uint64),
	}
}

func (bitmap *BitMap) Set(index uint64, bit bool) {
	base := index / bitlen
	offset := index % bitlen
	num := uint64(1) << offset
	if bit {
		bitmap.bm[base] = bitmap.bm[base] | num
	} else {
		//invert num
		num = ^num
		bitmap.bm[base] = bitmap.bm[base] & num
	}
}

func (bitmap *BitMap) Get(index uint64) bool {
	base := index / bitlen
	offset := index % bitlen
	num := uint64(1) << offset

	val, ok := bitmap.bm[base]
	if ok {
		return val & num != 0
	}
	return false
}
