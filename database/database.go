package database

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type ScheduleDatabase struct {
	orm *gorm.DB
}

func NewDatabase(user, pass, server, port, database string) *ScheduleDatabase {
	scheduleDatabase := &ScheduleDatabase{}

	err := scheduleDatabase.initDB(user, pass, server, port, database)
	if err != nil {
		log.Fatalf("Unable to initialize database connection due to: %s", err)
	}

	return scheduleDatabase
}

func (sdb *ScheduleDatabase) initDB(user, pass, server, port, database string) error {
	err := errors.New("")
	err = nil
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True",
		user, pass, server, port, database)
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