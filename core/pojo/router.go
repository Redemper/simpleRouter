package pojo

import (
	"gorm.io/gorm"
)

type Router struct {
	Name      string `gorm:"name",yaml:"name"`
	OriginUri string `gorm:"origin_uri",yaml:"origin-uri"`
	TargetUri string `gorm:"target_uri",yaml:"taget-uri"`
	Order     int    `gorm:"order",yaml:"order"`
	Enabled   bool   `gorm:"enabled",yaml:"enabled"`
	Filters   string `gorm:"filters",yaml:"filters"`
	gorm.Model
}
