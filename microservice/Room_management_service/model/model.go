package model

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
