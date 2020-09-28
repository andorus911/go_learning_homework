package main

import (
	"os"
	"strings"
)

func main() {

}

func ReadDir(dir string) (map[string]string, error) {
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	env := os.Environ()
	envMap := make(map[string]string)
	for _, v := range env {
		envVariable := strings.Split(v, "=")
		envMap[envVariable[0]] = envVariable[1]
	}
	return envMap, nil
}

//func RunCmd(cmd []string, env map[string]string) int {}
