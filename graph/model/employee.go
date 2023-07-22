package model

import (
	"io"
	"strconv"
	"time"

	"graphql-nosql/helper"

	"github.com/99designs/gqlgen/graphql"
)

type Employee struct {
	ID        string     `bson:"_id" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"text"`
	Age       int        `json:"done"`
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, helper.ServerError{Code: helper.TimeStampError}
}
