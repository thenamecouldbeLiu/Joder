package model

import (
	"time"
)

type PersonalInfo struct {
	UserId        string    `gorm:"column:user_id" binding:"required"`
	FacebookId    string    `gorm:"column:facebook_id"`
	GoogleId      string    `gorm:"column:google_id"`
	Gender        int       `gorm:"column:gender" binding:"required"`
	Email         string    `gorm:"column:email" binding:"required"`
	Birthday      time.Time `gorm:"column:birthday" binding:"required"` //string   `gorm:"type:date;column:birthday"`
	Name          string    `gorm:"column:name" binding:"required"`
	CreateAt      time.Time `gorm:"column:create_at"`
	UpdateAt      time.Time `gorm:"column:update_at"`
	UserType      int       `gorm:"column:user_type" binding:"required"`
	LastLoginTime time.Time `gorm:"column:last_login_time" binding:"required"`
}

func (PersonalInfo) TableName() string {
	return "tb_personal_info"
}

type UserIntro struct {
	UserId    string    `gorm:"column:user_id" binding:"required"`
	SelfIntro string    `gorm:"column:self_intro"`
	CreateAt  time.Time `gorm:"column:create_at"`
	UpdateAt  time.Time `gorm:"column:update_at"`
}

func (UserIntro) TableName() string {
	return "tb_user_intro"
}

type UserLocation struct {
	UserId    string    `gorm:"column:user_id" binding:"required"`
	Longitude float64   `gorm:"column:longitude"`
	Latitude  float64   `gorm:"column:latitude"`
	CreateAt  time.Time `gorm:"column:create_at"`
	UpdateAt  time.Time `gorm:"column:update_at"`
}

func (UserLocation) TableName() string {
	return "tb_user_loc"
}

type UserPhoto struct {
	UserId   string    `gorm:"column:user_id" binding:"required"`
	PhotoUrl string    `gorm:"column:photo_url"`
	CreateAt time.Time `gorm:"column:create_at"`
	UpdateAt time.Time `gorm:"column:update_at"`
}

func (UserPhoto) TableName() string {
	return "tb_user_photo"
}
