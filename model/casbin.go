package model

type CasbinRule struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Ptype string `gorm:"column:ptype;NOT NULL" json:"ptype"`
	V0    string `gorm:"column:v0;NOT NULL" json:"v0"`
	V1    string `gorm:"column:v1;NOT NULL" json:"v1"`
	V2    string `gorm:"column:v2;NOT NULL" json:"v2"`
	V3    string `gorm:"column:v3;NOT NULL" json:"v3"`
	V4    string `gorm:"column:v4;NOT NULL" json:"v4"`
	V5    string `gorm:"column:v5;NOT NULL" json:"v5"`
	V6    string `gorm:"column:v6;NOT NULL" json:"v6"`
	V7    string `gorm:"column:v7;NOT NULL" json:"v7"`
}
