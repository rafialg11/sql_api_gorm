package models

import "time"

type Book struct {
	ID        uint   `gorm : "primaryKey"`
	Title     string `gorm: "not null; unique; type:varchar(100)"`
	Author    string `gorm: "not null; unique; type:varchar(100)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
