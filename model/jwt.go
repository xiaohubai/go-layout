package model

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Uid        string `json:"uid"`
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	RoleId     string `json:"role_id"`
	RoleName   string `json:"role_name"`
	Birth      string `json:"brith"`
	State      string `json:"state"`
	BufferTime int64  `json:"buffer_time"`
	jwt.StandardClaims
}
