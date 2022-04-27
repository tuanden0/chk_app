package v1

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func uploadFiles(filename string, file multipart.File) (fPath string, err error) {

	workDir, _ := os.Getwd()
	fPath = filepath.Join(workDir, UPLOAD_DIR, filename)

	out, err := os.Create(fPath)
	if err != nil {
		return
	}
	defer out.Close()

	io.Copy(out, file)
	return
}
