package core

import (
	"io/ioutil"
	"os"
)

type fileAccess struct {
	readDir  func(dirname string) ([]os.FileInfo, error)
	readFile func(filename string) ([]byte, error)
}

var realFileAccess = fileAccess{
	readDir:  ioutil.ReadDir,
	readFile: ioutil.ReadFile}
