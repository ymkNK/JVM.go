package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func (z ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(z.absPath)
	if err!= nil {
		return nil, nil, err
	}
	defer reader.Close()
	for _, f:= range reader.File {
		if f.Name== className {
			rc,err :=f.Open()
			if err != nil {
				return nil,nil,err
			}
			data,err:=ioutil.ReadAll(rc)
			if err != nil{
				return nil, nil, err
			}
			rc.Close()
			return data,z,nil
		}
	}
	return nil, nil, errors.New("class not found: "+ className)
}

func (z ZipEntry) String() string {
	return z.absPath
}

func newZipEntry(path string) *ZipEntry {
	absPath,err:=filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}


