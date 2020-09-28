package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	
}

func ReadDir(dir string) (map[string]string, error) {
	filesFromDir := make(map[string]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		mode := info.Mode()
		if mode.IsRegular() {
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			filesFromDir[path] = string(b)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return filesFromDir, nil
}

//func RunCmd(cmd []string, env map[string]string) int {}
