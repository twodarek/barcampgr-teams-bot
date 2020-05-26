package database

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/twodarek/barcampgr-teams-bot/barcampgr"
	"log"
)

type ScheduleDatabase struct {
	orm *gorm.DB
	config barcampgr.Config
}

func NewDatabase(config barcampgr.Config) *ScheduleDatabase {
	scheduleDatabase := &ScheduleDatabase{
		config: config,
	}

	err := scheduleDatabase.initDB()
	if err != nil {
		log.Fatalf("Unable to initialize database connection due to: %s", err)
	}

	return scheduleDatabase
}

func (sdb *ScheduleDatabase) initDB() error {
	err := errors.New("")
	err = nil
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True",
		sdb.config.MySqlUser,
		sdb.config.MySqlPass,
		sdb.config.MySqlServer,
		sdb.config.MySqlPort,
		sdb.config.MySqlDatabase)
	sdb.orm, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}
	return nil
}

func (sdb *ScheduleDatabase) MigrateDB() error {
	sdb.orm.AutoMigrate(&DBScheduleRoom{}, &DBScheduleTime{}, &DBScheduleSession{})
	return nil
}