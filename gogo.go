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

	err = errors.New("src dirctory no found")
	return
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {

	orgcwd, err := os.Getwd()
	exitOnError(err)

	cwd := orgcwd

	err = os.Chdir("src")
	if err == nil {
		cwd, err = os.Getwd()
		exitOnError(err)

		err = os.Chdir(orgcwd)
		exitOnError(err)
	}

	fmt.Println(cwd)
	gopath, err := getGOPATH(cwd)
	exitOnError(err)

	fmt.Println(gopath)
	fmt.Println(os.Args)

	err = runproc("go", gopath, os.Args[1:])
	exitOnError(err)
}
