package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please check the env directory and application to run.")
	}
	t, _ := ReadDir(os.Args[1])
	c := RunCmd(os.Args[2:], t)
	os.Exit(c)
}

// ReadDir reads directory files and their content
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
			filesFromDir[info.Name()] = string(b)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return filesFromDir, nil
}

// RunCmd runs a command with special env and returns exit code of the command
func RunCmd(cmd []string, env map[string]string) int {
	cmdToRun := exec.Command(cmd[0], cmd[1:]...)

	cmdToRun.Stdout = os.Stdout
	cmdToRun.Stderr = os.Stderr
	cmdToRun.Stdin = os.Stdin

	envSlice := make([]string, len(env))
	i := 0
	for key, v := range env {
		envSlice[i] = key + "=" + v
		i++
	}
	cmdToRun.Env = envSlice // or append

	if err := cmdToRun.Run(); err != nil {
		log.Println("Command finished with error:", err)
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}
	return 0
}
