package ioutils

import (
	"io"
	"os"
	"io/ioutil"
)

// Terminate is drop all and close
func Terminate(r io.ReadCloser) error {
	_, err := io.Copy(ioutil.Discard, r)
	if err1 := r.Close(); err == nil {
		err = err1
	}
	return err
}

func ReadDir(dir string) ([]os.FileInfo, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func ReadDirNames(dir string) ([]string, error) {
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	list, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	return list, nil
}

