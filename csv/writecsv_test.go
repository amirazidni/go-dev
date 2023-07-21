package gocsv

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestWriteCSV(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func TestWriteAllCSV(t *testing.T) {
	records := [][]string{
		{"first_name", "last_name", "username", "full_name"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

// =============================================================
type User struct {
	FirstName string
	LastName  string
	UserName  string
	FullName  string
}

func TestWrite2FileCSV(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	// call csv builder
	if err := CreateFile(); err != nil {
		fmt.Println(err)
	}

	// download file
	writer.Header().Add("Content-Disposition", "attachment; filename=\"download.csv\"")
	http.ServeFile(writer, request, "./download.csv")

	// remove csv
	if err := os.Remove("download.csv"); err != nil {
		fmt.Println(err)
	}
}

func CreateFile() error {
	records := []User{
		{"first_name", "last_name", "username", "full_name"},
		{"Rob", "Pike", "rob", ""},
		{"Ken", "Thompson", "ken", ""},
		{"Robert", "Griesemer", "gri", ""},
	}

	// create csv file
	file, err := os.Create("download.csv")
	defer file.Close()
	if err != nil {
		return err
	}

	// write inside
	w := csv.NewWriter(file)
	defer w.Flush()

	// use write
	for _, record := range records {
		row := []string{record.FirstName, record.LastName, record.UserName, record.FullName}
		if err := w.Write(row); err != nil {
			return err
		}
	}

	// use writeAll
	var data [][]string
	for _, record := range records {
		row := []string{record.FirstName, record.LastName, record.UserName, record.FullName}
		data = append(data, row)
	}
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
		return err
	}
	return nil
}
