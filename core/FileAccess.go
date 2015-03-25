package core

import (
	"io/ioutil"
	"os"
)

type fileAccess struct {
	readDir func(dirname string) ([]os.FileInfo, error)
}

var realFileAccess = fileAccess{
	readDir: ioutil.ReadDir}
