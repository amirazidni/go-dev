package test

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// will import file automatic
//go :embed version.txt
var version string

//go :embed bangkit.png
var logo []byte

//go :embed files/test_files.txt
//go :embed files/test_files2.txt
//go :embed files/test_files3.txt
var files embed.FS

// use regex
//go :embed files/*.txt
var path embed.FS

func TestString(t *testing.T) {
	fmt.Println(version)
}

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("newLogo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(logo)
}

func TestMultipleFiles(t *testing.T) {
	file1, _ := files.ReadFile("files/test_files.txt")
	fmt.Println(string(file1))

	file2, _ := files.ReadFile("files/test_files2.txt")
	fmt.Println(string(file2))

	file3, _ := files.ReadFile("files/test_files3.txt")
	fmt.Println(string(file3))
}

func TestPathMatcher(t *testing.T) {
	// read directory
	dirEntries, _ := path.ReadDir("files")

	for _, de := range dirEntries {
		// if de is a file
		if !de.IsDir() {
			fmt.Println(de.Name())
			file, _ := path.ReadFile("files/" + de.Name())
			fmt.Println(string(file))
		}

	}
}
