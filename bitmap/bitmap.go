// Simple Bitmap using a map
package bitmap

// Simple BitMap
//
// only stores chunks with true values
// sorta
type BitMap struct {
	bm map[uint64]uint64
}

const bitlen uint64 = 64

// Simple constructor for BitMap
func NewBitMap() *BitMap {
	return &BitMap{
		bm: make(map[uint64]uint64),
	}
}

// Sets bit at index to bit value
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

// Gets bit value at index
//
// Returns false if never set.
func (bitmap *BitMap) Get(index uint64) bool {
	base := index / bitlen
	offset := index % bitlen
	num := uint64(1) << offset

	val, ok := bitmap.bm[base]
	if ok {
		return val&num != 0
	}
	return false
}
