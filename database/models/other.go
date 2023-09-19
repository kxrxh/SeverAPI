package models

import "gorm.io/gorm"

type City struct {
	gorm.Model        // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Name       string `gorm:"column:name;not null"`
}

type State struct {
	gorm.Model        // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Name       string `gorm:"column:name;not null"`
}

type Route struct {
	gorm.Model        // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Number     string `gorm:"column:number;not null"`
}

type Transport struct {
	UID         string `gorm:"primaryKey;column:uid"`
	Number      string `gorm:"column:number;not null"`
	IsAvailable bool   `gorm:"column:is_available;not null"`
}

type ClientBenefit struct {
	gorm.Model      // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	ClientID   uint `gorm:"column:client_id;not null;primaryKey"`
	BenefitID  uint `gorm:"column:benefit_id;not null;primaryKey"`
	Client     Client
	Benefit    Benefit
}

type ClientPromotion struct {
	gorm.Model       // GORM will add ID, CreatedAt, UpdatedAt, and DeletedAt fields
	ClientID    uint `gorm:"column:client_id;not null;primaryKey"`
	PromotionID uint `gorm:"column:promotion_id;not null;primaryKey"`
	Client      Client
	Promotion   Promotion
}
