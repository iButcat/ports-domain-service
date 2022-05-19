package utils

import (
	"mime/multipart"
	"net/http"
)

func ParseMultipartAndReturnFile(r *http.Request) (multipart.File, *multipart.FileHeader, error) {
	const maxMemory = 32 << 20 // 32 MB
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		return nil, nil, err
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return nil, nil, err
	}
	return file, handler, nil
}
