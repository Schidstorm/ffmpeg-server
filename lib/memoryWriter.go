package lib

type MemoryWriter struct {
	buffer []byte
}

func CreateMemoryWriter() *MemoryWriter {
	return &MemoryWriter{buffer: make([]byte, 0)}
}

func (w *MemoryWriter) Write(data []byte) (int, error) {
	w.buffer = append(w.buffer, data...)
	return len(data), nil
}

func (w *MemoryWriter) Buffer() []byte {
	return w.buffer
}