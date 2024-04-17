package types

// HeightQueue is a simple type alias for a (buffered) channel of block heights.
type HeightQueue chan int64

func NewQueue(size int) HeightQueue {
	return make(chan int64, size)
}

// RetriesCount is a simple type alias for a map of block heights to the number of retries.
type RetriesCount map[int64]int64

func NewRetriesCount() RetriesCount {
	return make(map[int64]int64)
}

func (r RetriesCount) Increment(height int64) {
	if r == nil {
		return
	}
	r[height]++
}

func (r RetriesCount) Get(height int64) int64 {
	if r == nil {
		return 0
	}
	return r[height]
}
