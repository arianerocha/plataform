package db

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/krakenlab/plataform/internal/configs"

	// Gorm dialects
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var singleConn *gorm.DB

// Connection aready configurated to db
func Connection() *gorm.DB {
	if singleConn == nil {
		singleConn = tryConnect()
	}

	return singleConn
}

func tryConnect() *gorm.DB {
	db, err := gorm.Open(configs.Db().Adapter(), configs.Db().URL())
	if err != nil {
		log.Print(err)
		time.Sleep(5 * time.Second)
		return tryConnect()
	}

	return db.LogMode(true)
}
