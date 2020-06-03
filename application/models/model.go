package models

import "time"

type Model struct {
	ID        uint       `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	CreatedAt time.Time  `sql:"timestamp;not null;default:'2006-02-03 15:04:05'"`
	UpdatedAt time.Time  `sql:"timestamp;not null;default:'2006-02-03 15:04:05'"`
	DeletedAt *time.Time `sql:"timestamp;not null;default:'2006-02-03 15:04:05'"`
}
