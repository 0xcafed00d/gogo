package main

import (
	//"github.com/simulatedsimian/assert"
	"testing"
)

func TestGOPATH(t *testing.T) {
	getGOPATH("/home/lmw/src/test/wibble/")
	getGOPATH("c:/home/lmw/src/test/wibble/")
}
