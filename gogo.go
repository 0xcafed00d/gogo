package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

	src := filepath.Join(string(os.PathSeparator), "src")

	for {
		if strings.HasSuffix(cwd, src) {
			gopath = cwd[:len(cwd)-len(src)]
			return
		}

		sepIdx := strings.LastIndex(cwd, string(os.PathSeparator))
		if sepIdx == -1 {
			break
		}
		cwd = cwd[:sepIdx]
	}

	err = errors.New("Error: src directory not found")
	return
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	path, err := os.Getwd()
	exitOnError(err)

	finfo, err := os.Stat("src")
	if err == nil && finfo.IsDir() {
		path = filepath.Join(path, "src")
	}

	gopath, err := getGOPATH(path)
	exitOnError(err)

	err = runproc("go", gopath, os.Args[1:])
	exitOnError(err)
}
