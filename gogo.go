package main

import (
	"errors"
	"fmt"
	//	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	//	"time"
)

// execs a process with the supplied GOPATH
func runproc(proc, gopath string, args []string) error {

	cmd := exec.Command(proc, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	for _, evar := range os.Environ() {
		if strings.HasPrefix(evar, "GOPATH=") {
			cmd.Env = append(cmd.Env, "GOPATH="+gopath)
		} else {
			cmd.Env = append(cmd.Env, evar)
		}
	}

	return cmd.Run()

}

func getGOPATH(cwd string) (gopath string, err error) {
	cderr := os.Chdir("src")
	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	cwdDirs := filepath.SplitList(cwd)

	os.PathSeparator
	errors.New("src directory not found in current path")

	return
}

func main() {
	fmt.Println(os.Getwd())

	os.Chdir("src")
	fmt.Println(os.Getwd())

}
