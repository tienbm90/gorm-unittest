package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	//ID    uint   `gorm:"primary_key"`
	Code  string `gorm:"column:code" json:"code"`
	Price uint	 `gorm:"column:price" json:"price"`
}

func (p *Product) TableName() string {
	return "products"
}
