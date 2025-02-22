package model

import (
	"user_interface/userpb"

	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Data userpb.User `json:"username"`
	jwt.StandardClaims
}

type User struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Age        int64  `json:"age"`
	Email      string `json:"email"`
	Contact    string `json:"contact"`
	Manager    bool   `json:"manager"`
	Department string `json:"department"`
}

type User_delete struct {
	Name string `json:"name"`
}

//////////////////////////////////////////////////

type Room struct {
	Room_number     int64    `json:"room_number"`
	Price           float32  `json:"price"`
	Features        []string `json:"features"`
	OccupencyStatus bool     `json:"Occupency"`
}

type Room_feature_update struct {
	Room_numer int64    `json:"room_number"`
	Features   []string `json:"features"`
}

type Room_price_update struct {
	Room_number int64   `json:"room_number"`
	Price       float32 `json:"price"`
}

type Room_Delete struct {
	Room_number int64 `json:"room_number"`
}

////////////////////////////////////////

type Customer struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}

type Total struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
	OTP     string `json:"otp"`
}

type Claims_total struct {
	Data Total `json:"Data"`
	jwt.StandardClaims
}
