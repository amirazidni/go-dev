package model

import (
	"go-generic/pkg/utils"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"gorm.io/gorm"
)

type Todo struct {
	ID        string          `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	Text      string          `json:"text"`
	Done      bool            `json:"done"`
	UserID    string          `json:"user_id"`
	User      *User           `json:"user" gorm:"ForeignKey:UserID"`
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
	return time.Time{}, utils.TimeStampError
}

func MarshalNulltime(d gorm.DeletedAt) graphql.Marshaler {
	t := d.Time
	timestamp := t.Unix() * 1000
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalNulltime(v interface{}) (gorm.DeletedAt, error) {
	val := gorm.DeletedAt{}
	if tmpStr, ok := v.(int); ok {
		val.Time = time.Unix(int64(tmpStr), 0)
		return val, nil
	}
	return val, utils.TimeStampError
}
