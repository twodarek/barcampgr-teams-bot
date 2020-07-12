package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DBScheduleTime struct {
	gorm.Model
	Start string
	End string
	Displayable bool `gorm:"default:false"`
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

func (s DBScheduleSession) ToString() string {
	return fmt.Sprintf("Title: %s, Speaker: %s, Start Time: %s, Room: %s", s.Title, s.Speaker, s.Time.Start, s.Room.Name)
}
