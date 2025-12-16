package db

// ChunkSlice splits input into chunks of size chunkSize.
func chunkSlice[T any](slice []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		chunkSize = 1000
	}

	var chunks [][]T
	for start := 0; start < len(slice); start += chunkSize {
		end := start + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[start:end])
	}
	return chunks
}
