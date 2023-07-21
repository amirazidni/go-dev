package golang_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

// will import file automatic
//go:embed version.txt
var Version string

//go:embed bangkit.png
var Logo []byte

// use regex
//go:embed files/*.txt
var Path embed.FS

func main() {
	fmt.Println(Version)
	err := ioutil.WriteFile("newLogo.png", Logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	// read directory
	dirEntries, _ := Path.ReadDir("files")

	for _, de := range dirEntries {
		// if de is a file
		if !de.IsDir() {
			fmt.Println(de.Name())
			file, _ := Path.ReadFile("files/" + de.Name())
			fmt.Println(string(file))
		}

	}
}
