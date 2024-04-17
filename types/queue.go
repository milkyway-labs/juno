package types

// HeightQueue is a simple type alias for a (buffered) channel of block heights.
type HeightQueue chan int64

func NewQueue(size int) HeightQueue {
	return make(chan int64, size)
}

// RetriesCount is a simple type alias for a map of block heights to the number of retries.
type RetriesCount struct {
	maxRetries int64
	count      map[int64]int64
}

func NewRetriesCount(maxRetries int64) *RetriesCount {
	return &RetriesCount{
		maxRetries: maxRetries,
		count:      make(map[int64]int64),
	}
}

func (r *RetriesCount) Increment(height int64) {
	if r == nil {
		return
	}
	r.count[height]++
}

func (r *RetriesCount) HasReachedMax(height int64) bool {
	if r == nil {
		return false
	}
	return r.maxRetries != -1 && r.count[height] >= r.maxRetries
}

func (r *RetriesCount) Get(height int64) int64 {
	if r == nil {
		return 0
	}
	return r.count[height]
}
