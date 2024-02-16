package storage

import (
	"gorm.io/gorm"
)

type UserDAO struct {
	gorm.Model
	UserID              uint64
	FontSize            string
	DarkMode            bool
	ReceiveNotification bool
	Language            string
	Notifications       bool
	Nickname            string
	Avatar              string
	Signature           string
	Location            string
	Age                 int
	Gender              string
	Tags                string
}

func (u UserDAO) TableName() string {
	return "user"
}

type DeviceDAO struct {
	gorm.Model
	DeviceID   uint64
	OS         string
	AppVersion string
	Type       string
	Mode       string
	UserID     uint64
}

func (d DeviceDAO) TableName() string {
	return "device"
}
