package models

import (
	"time"

	"gorm.io/gorm"
)

type Conductor struct {
	gorm.Model        // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	FirstName  string `gorm:"column:first_name;not null"`
	LastName   string `gorm:"column:last_name;not null"`
	MiddleName string `gorm:"column:middle_name"`
	Sex        string `gorm:"column:sex;type:char(1);check:sex IN ('М', 'Ж');not null"`
	Email      string `gorm:"column:email;not null"`
	Phone      string `gorm:"column:phone;not null"`
	DeviceUID  string `gorm:"column:device_uid;not null;unique"`
	StateID    uint   `gorm:"column:state_id;references:id"`
	State      State  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming State is a related model
}

type ConductorCard struct {
	gorm.Model           // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	CardUID    string    `gorm:"column:card_uid;not null;unique"`
	StateID    uint      `gorm:"column:state_id;references:id"`
	State      State     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming State is a related model
	TimeOpen   time.Time `gorm:"column:time_open;not null"`
}

type Manager struct {
	gorm.Model           // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	FirstName  string    `gorm:"column:first_name;not null"`
	LastName   string    `gorm:"column:last_name;not null"`
	MiddleName string    `gorm:"column:middle_name"`
	Sex        string    `gorm:"column:sex;type:char(1);check:sex IN ('М', 'Ж');not null"`
	Email      string    `gorm:"column:email;not null"`
	Phone      string    `gorm:"column:phone;not null"`
	Birth      time.Time `gorm:"column:birth;not null"`
	StateID    uint      `gorm:"column:state_id;references:id"`
	State      State     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming State is a related model
}

type Shift struct {
	gorm.Model            // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Action      string    `gorm:"column:action;not null;type:char(3);check:action IN ('отк', 'зак')"`
	Time        time.Time `gorm:"column:time;not null"`
	ConductorID uint      `gorm:"column:conductor_id;references:id"`
	Conductor   Conductor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming Conductor is a related model
	RouteID     uint      `gorm:"column:route_id;references:id"`
	Route       Route     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming Route is a related model
	TransportID uint      `gorm:"column:transport_id;references:id"`
	Transport   Transport `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming Transport is a related model
}

