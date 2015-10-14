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

func Split(s, charset string) []string {
	res := []string{}
	tokenStart := -1

	for i, r := range s {
		if strings.ContainsRune(charset, r) {
			if tokenStart != -1 {
				res = append(res, s[tokenStart:i])
				tokenStart = -1
			}
		} else {
			if tokenStart == -1 {
				tokenStart = i
			}
		}
	}
	if tokenStart != -1 {
		res = append(res, s[tokenStart:])
	}
	return res
}

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

	cwdDirs := Split(cwd, string(os.PathSeparator))

	for i := len(cwdDirs) - 1; i > -1; i-- {
		fmt.Println(i, cwdDirs[i])
		if cwdDirs[i] == "src" {
			for n := 0; n < i; n++ {
				gopath = filepath.Join(cwdDirs...)
				fmt.Println(gopath)
				return
			}
		}
	}

	fmt.Println(cwdDirs)
	err = errors.New("src dirctory no found")
	return
}

func main() {
	fmt.Println(os.Getwd())

	os.Chdir("src")
	fmt.Println(os.Getwd())

}
