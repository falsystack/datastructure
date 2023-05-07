package buffer

type RingBuffer[T any] struct {
	buf     []T
	readPt  int
	writePt int
}

func (r *RingBuffer[T]) Read(count int) []T {
	readCnt := min(count, r.Readable())
	rst := make([]T, readCnt)
	copy(rst, r.buf[r.readPt:])
	r.readPt += readCnt
	return rst
}

func (r *RingBuffer[T]) Writable() int {
	if r.writePt < r.readPt {
		return len(r.buf) - r.readPt + r.writePt
	}
	return len(r.buf) - r.Readable()
}

func (r *RingBuffer[T]) Readable() int {
	return r.writePt - r.readPt
}

func (r *RingBuffer[T]) Write(data []T) int {
	if len(data) == 0 {
		return 0
	}
	if r.writePt >= r.readPt {
		writableToEnd := len(r.buf) - r.writePt
		writed := min(writableToEnd, len(data))
		copy(r.buf[r.writePt:], data[:writed])
		r.writePt += writed
		r.writePt %= len(r.buf) // 重要

		remain := len(data) - writed
		if remain > 0 {
			return writed + r.Write(data[writed:])
		}
		return writed
	}
	availableCnt := min(r.Writable(), len(data))
	if availableCnt == 0 {
		return availableCnt
	}
	copy(r.buf[r.writePt:], data[:availableCnt])
	r.writePt += availableCnt
	return availableCnt
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, size),
	}
}
