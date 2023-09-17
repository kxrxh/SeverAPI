package models

import (
	"time"

	"gorm.io/gorm"
)

type Benefit struct {
	gorm.Model           // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Value      string    `gorm:"column:value;not null"`
	Name       string    `gorm:"column:name;not null;unique"`
	BeginTime  time.Time `gorm:"column:begin_time;not null"`
	EndTime    time.Time `gorm:"column:end_time;not null"`
}

type Partner struct {
	gorm.Model        // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Name       string `gorm:"column:name;not null"`
	Email      string `gorm:"column:email;not null"`
	Password   string `gorm:"column:password;not null"`
	StateID    uint   `gorm:"column:state_id;references:id"`
	State      State  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming State is a related model
}

type Promotion struct {
	gorm.Model           // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Name       string    `gorm:"column:name;not null"`
	BeginTime  time.Time `gorm:"column:begin_time;not null"`
	EndTime    time.Time `gorm:"column:end_time;not null"`
	PartnerID  uint      `gorm:"column:partner_id;references:id"`
	Partner    Partner   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Assuming Partner is a related model
}
