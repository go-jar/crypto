package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func Md5String(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Md5File(file io.Reader) (string, error) {
	m := md5.New()
	if _, err := io.Copy(m, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", m.Sum(nil)), nil
}

func Md5FileByPath(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	return Md5File(f)
}
