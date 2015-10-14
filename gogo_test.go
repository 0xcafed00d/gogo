package main

import (
	"github.com/simulatedsimian/assert"
	"testing"
)

func TestGOPATH(t *testing.T) {

	pack := assert.Pack

	assert.HasError(t, pack(getGOPATH("")))
	assert.HasError(t, pack(getGOPATH("/")))
	assert.HasError(t, pack(getGOPATH("c:/home/lmw")))

	assert.Equal(t, pack(getGOPATH("/home/lmw/src/test/wibble")), pack("/home/lmw", nil))
	assert.Equal(t, pack(getGOPATH("/home/lmw/src/src")), pack("/home/lmw/src", nil))

	assert.Equal(t, pack(getGOPATH("c:/home/lmw/src/test/wibble")), pack("c:/home/lmw", nil))
}
