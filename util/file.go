package util

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/tal-tech/go-zero/tools/goctl/util"
)

func CreateFiles(modelList map[string]string, dir string) error {
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	err = util.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	for fileName, code := range modelList {
		filename := filepath.Join(dirAbs, fileName)
		if util.FileExists(filename) {
			logrus.Warnf("%s already exists, ignored.", fileName)
			continue
		}
		err = ioutil.WriteFile(filename, []byte(code), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
