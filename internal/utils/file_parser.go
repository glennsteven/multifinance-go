package utils

import (
	"bytes"
	"errors"
	"fmt"
	"multifinance-go/internal/presentations"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func MultipartFormFile(r *http.Request, field string, maxFileSize int64, extension []string) (*presentations.File, error) {
	f, h, err := r.FormFile(field)
	if errors.Is(err, http.ErrMissingFile) {
		return nil, fmt.Errorf("the %s image field required", field)
	}

	if err != nil {
		return nil, fmt.Errorf("parse form file %s err %v", field, err)
	}

	var buff bytes.Buffer
	size, err := buff.ReadFrom(f)
	if err != nil {
		return nil, nil
	}

	if maxFileSize != 0 && size > maxFileSize {
		return nil, fmt.Errorf("the %s image file size %s is too large, max allow is %s", field, strconv.FormatInt(h.Size, 10), strconv.FormatInt(maxFileSize, 10))
	}

	ct := ExtractFileExtension(buff.Bytes())
	ext := FileExtension(ct)

	if len(extension) != 0 && !ValidFileExtension(ext, extension) {
		return nil, fmt.Errorf("the %s image file extension  .%s not allowed, request original file content type is %s, only allow: %s", field, ext, ct, strings.Join(extension, ", "))
	}

	fileExt := GetFileNameExtension(h.Filename)
	if len(extension) != 0 && !ValidFileExtension(fileExt, extension) && fileExt != "" {
		return nil, fmt.Errorf(fmt.Sprintf("the %s image file extension .%s not allowed, only know: %s", field, fileExt, strings.Join(extension, ", ")))
	}

	result := &presentations.File{
		Filename:    h.Filename,
		Ext:         fileExt,
		ContentType: ct,
		Size:        size,
		Buffer:      &buff,
	}

	return result, nil
}

func ExtractFileExtension(data []byte) string {
	return http.DetectContentType(data)
}

func FileExtension(input string) string {
	if len(input) < 1 {
		return input
	}
	return input[strings.IndexByte(input, '/')+1:]
}

func ValidFileExtension(ext string, extension []string) bool {
	return InArray(ext, extension)
}

func GetFileNameExtension(input string) string {
	ex := filepath.Ext(input)
	if ex == "" {
		return ex
	}

	return strings.ToLower(ex[1:len(ex)])
}
