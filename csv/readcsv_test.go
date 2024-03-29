package gocsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestReadCSV(t *testing.T) {
	in := `
first_name,last_name,username
Rob,Pike,rob
Ken,Thompson,ken
Robert,Griesemer,gri`
	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func TestReadCSVfile(t *testing.T) {
	reader, err := os.Open("download.csv")
	defer reader.Close()
	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(reader)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
