package types

// BlockData contains the data of a block that is used to be stored in the queue.
type BlockData struct {
	Height     int64
	RetryCount int64
}

func NewBlockData(height int64) BlockData {
	return BlockData{
		Height: height,
	}
}

func (b BlockData) HasReachedMaxRetries(maxRetries int64) bool {
	return maxRetries != -1 && b.RetryCount >= maxRetries
}

func (b BlockData) IncrementRetryCount() BlockData {
	return BlockData{
		Height:     b.Height,
		RetryCount: b.RetryCount + 1,
	}
}

// HeightQueue is a simple type alias for a (buffered) channel of block heights.
type HeightQueue chan BlockData

func NewQueue(size int) HeightQueue {
	return make(chan BlockData, size)
}
