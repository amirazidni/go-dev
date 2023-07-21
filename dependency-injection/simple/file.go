package simple

import "fmt"

type File struct {
	Name string
}

func (f *File) Close() {
	fmt.Println("Close file", f.Name)
}

// clean up function
func NewFile(name string) (*File, func()) { // function akhir sebagai cleanup
	file := &File{Name: name}
	return file, func() {
		file.Close()
	}
}
