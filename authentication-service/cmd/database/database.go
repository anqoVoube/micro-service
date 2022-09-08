package database

import (
	"authentication/config"
	"sync"

	"github.com/go-pg/pg/v10"
)

type DB struct {
	*pg.DB
}

var (
	once sync.Once
	db   DB
)

func Get() *DB {
	once.Do(func() {
		configData := *config.New()
		//address := fmt.Sprintf("%s:%s", "postgres", "5432")
		options := &pg.Options{
			User:     configData.Postgres.User,
			Password: configData.Postgres.Password,
			Addr:     configData.Addr(),
			Database: configData.Postgres.Database,
			PoolSize: configData.Postgres.Poolsize,
		}
		pgDB := pg.Connect(options)
		db.DB = pgDB
	})
	return &db
}
