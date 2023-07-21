package simple

import "fmt"

type Connection struct {
	*File
}

func (c *Connection) Close() {
	fmt.Println("Close connection", c.File.Name)
}

// clean up function
func NewConnection(file *File) (*Connection, func()) { // closure
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}
