package models

import (
	"time"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model            // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	FirstName  string     `gorm:"column:first_name;not null"`
	LastName   string     `gorm:"column:last_name;not null"`
	MiddleName string     `gorm:"column:middle_name"`
	Phone      string     `gorm:"column:phone;not null;unique"`
	Sex        string     `gorm:"column:sex;type:char(1);check:sex IN ('М', 'Ж');not null"`
	Email      string     `gorm:"column:email;not null"`
	CityID     uint       `gorm:"column:city_id;references:id"`
	City       City       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming City is a related model
	Birth      time.Time  `gorm:"column:birth;not null"`
	StateID    uint       `gorm:"column:state_id;references:id"`
	State      State      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming State is a related model
	CardID     uint       `gorm:"column:card_id;references:id"`
	Card       ClientCard `gorm:"foreignKey:CardID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Snils      string     `gorm:"column:snils;not null;unique"`
}

type ClientCard struct {
	gorm.Model           // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	UID        string    `gorm:"column:uid;not null;unique"`
	TimeOpen   time.Time `gorm:"column:time_open;not null"`
	TimeClose  time.Time `gorm:"column:time_close;not null"`
	StateID    uint      `gorm:"column:state_id;references:id"`
	State      State     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming State is a related model
	ClientID   uint      `gorm:"column:client_id;references:id"`
}

type Operation struct {
	gorm.Model           // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Time       time.Time `gorm:"column:time;not null"`
	CordX      *int      `gorm:"column:cord_x"` // Use pointer type to represent nullable columns
	CordY      *int      `gorm:"column:cord_y"` // Use pointer type to represent nullable columns
	DeviceUID  string    `gorm:"column:device_uid;not null"`
	Benefits   bool      `gorm:"column:benefits;not null"`
}

type Benefit struct {
	gorm.Model           // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Value      string    `gorm:"column:value;not null"`
	Name       string    `gorm:"column:name;not null;unique"`
	BeginTime  time.Time `gorm:"column:begin_time;not null"`
	EndTime    time.Time `gorm:"column:end_time;not null"`
}
