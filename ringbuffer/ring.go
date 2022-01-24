package ringbuffer

type RingBuffer[T any] struct {
	buff  []T
	size  int
	start int
	max   int
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buff:  make([]T, size),
		size:  0,
		max:   size,
		start: 0,
	}
}

//len?
func (rb *RingBuffer[T]) Size() int {
	return rb.size
}

func (rb *RingBuffer[T]) Add(item T) {
	index := (rb.start + rb.size) % rb.max
	rb.buff[index] = item
	if rb.size < rb.max {
		rb.size++
	} else {
		rb.start++
	}
}

func (rb *RingBuffer[T]) Get(index int) T {
	if index >= rb.size {
		panic("out of range")
	}
	real := (rb.start + index) % rb.max
	return rb.buff[real]
}
