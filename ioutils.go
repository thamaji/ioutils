package ioutils

import (
	"io"
	"io/ioutil"
	"os"
)

// Terminate is drop all and close
func Terminate(r io.ReadCloser) error {
	_, err := io.Copy(ioutil.Discard, r)
	if err1 := r.Close(); err == nil {
		err = err1
	}
	return err
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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

func ReadDirOrEmpty(dir string) []os.FileInfo {
	list, _ := ReadDir(dir)
	return list
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

func ReadDirNamesOrEmpty(dir string) []string {
	list, _ := ReadDirNames(dir)
	return list
}
