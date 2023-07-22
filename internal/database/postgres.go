package database

import (
	"go-generic/pkg/utils"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *PostgreSQL
)

type PostgreSQL struct {
	Gorm *gorm.DB
}

func DbOrm() *gorm.DB {
	var err error
	once.Do(func() {
		db = &PostgreSQL{}
		db.Gorm, err = gorm.Open(
			postgres.Open(
				"host = localhost "+
					"user = root "+
					"password = root "+
					"dbname = go_generic "+
					"port = 5432 "+
					"sslmode = disable "+
					"TimeZone = Asia/Jakarta",
			),
			&gorm.Config{},
		)
		utils.ErrorHandler(err, "failed to open postgresql connection")
	})
	return db.Gorm
}
