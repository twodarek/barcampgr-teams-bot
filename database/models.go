package database

import "github.com/jinzhu/gorm"

type DBScheduleTime struct {
	gorm.Model
	Start string
	End string
}

type DBScheduleRoom struct {
	gorm.Model
	Name string
}

type DBScheduleSession struct {
	gorm.Model
	Time DBScheduleTime
	Room DBScheduleRoom
	Updater string
	Title string
	Speaker string
}
