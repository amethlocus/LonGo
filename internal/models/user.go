package models

import "time"

type User struct {
	Username   string    `form:"username" json:"username" gorm:"primaryKey;size:30"`
	Email      string    `form:"email" json:"email" gorm:"unique;not null"`
	Password   string    `form:"password" json:"password" gorm:"not null"`
	Name       string    `form:"name" json:"name"`
	Lastname   string    `form:"lastname" json:"lastname"`
	CreatedAt  time.Time `form:"created_at" json:"created_at" gorm:"not null;default:now()"`
	LastUpdate time.Time `form:"last_update" json:"last_update" gorm:"not null;default:now()"`
}
