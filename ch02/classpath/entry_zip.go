package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	/**
	 * zip或jar文件的绝对路径
	 */
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	zipReadCloser, err := zip.OpenReader(self.absPath)
	if err != nil {
		panic(err)
	}
	defer zipReadCloser.Close()
	for _, f := range zipReadCloser.File {
		if f.Name == className {
			fileReadCloser, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer fileReadCloser.Close()
			data, err := ioutil.ReadAll(fileReadCloser)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
