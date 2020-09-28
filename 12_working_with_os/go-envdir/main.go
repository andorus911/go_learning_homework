package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

func RunCmd(cmd []string, env map[string]string) int {
	cmdToRun := exec.Command(cmd[0], cmd[1:]...)

	envSlice := make([]string, len(env))
	i := 0
	for key, v := range env {
		envSlice[i] = key + "=" + v
		i++
	}
	cmdToRun.Env = envSlice

	err := cmdToRun.Start()
	if err != nil {
		log.Println("Command started with error:", err)
	}
	err = cmdToRun.Wait()
	if err != nil {
		log.Println("Command finished with error:", err)
	}
	return 0 // TO DO
}
