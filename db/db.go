package db

import (
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ms-go-initial/config"
	"strings"
	"time"
)

var db *gorm.DB

func InitDb() {
	host := strings.Split(config.Conf.DBUrl, ":")[0]
	port := strings.Split(config.Conf.DBUrl, ":")[1]
	user := config.Conf.DBUser
	pass := config.Conf.DBPass
	name := config.Conf.DBName

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Baku", host, port, user, pass, name)

	tempDb, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Error("ActionLog.DbInit.error ", err)
	}
	pqDB, err := tempDb.DB()
	if err != nil {
		log.Printf("ActionLog.DbInit.error  %v", err)
	}
	pqDB.SetMaxIdleConns(2)                   //Maximum number of free connections
	pqDB.SetMaxOpenConns(5)                   //maximum connection
	pqDB.SetConnMaxLifetime(time.Minute * 10) //Set connection idle timeout
	db = tempDb
}

func GetDb() *gorm.DB { return db }
