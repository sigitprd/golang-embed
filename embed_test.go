package golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed image.png
var image []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("image_new.png", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
