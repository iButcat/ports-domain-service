package utils

import (
	"io"
	"mime/multipart"
)

// should use json decode with Next method to handle it as a steam.
func ReadFileByChunks(file multipart.File, chunkSize int64) ([]byte, error) {
	buffer := make([]byte, chunkSize)
	var chunks []byte
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			chunks = append(chunks, buffer[:n]...)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return chunks, nil
}
