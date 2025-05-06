package gwm_gorm

import (
	"fmt"
	"sync"

	gwm_app "github.com/slhmy/go-webmods/app"
	"github.com/slhmy/go-webmods/internal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	initMutx sync.Mutex
	db       *gorm.DB
)

func GetDB() *gorm.DB {
	if db == nil {
		initMutx.Lock()
		defer initMutx.Unlock()
		if db != nil {
			return db
		}

		driver := gwm_app.Config().GetString(internal.ConfigKeyDatabaseDriver)
		switch driver {
		case "postgres":
			db, err := openPostgres()
			if err != nil {
				panic(err)
			}
			return db
		default:
			panic(fmt.Sprintf("unsupported database driver: %s", driver))
		}
	}
	return db
}

func openPostgres() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		gwm_app.Config().GetString(internal.ConfigKeyDatabaseHost),
		gwm_app.Config().GetString(internal.ConfigKeyDatabasePort),
		gwm_app.Config().GetString(internal.ConfigKeyDatabaseUsername),
		gwm_app.Config().GetString(internal.ConfigKeyDatabaseName),
		gwm_app.Config().GetString(internal.ConfigKeyDatabasePassword),
		gwm_app.Config().GetString(internal.ConfigKeyDatabaseSSLMode),
	)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
