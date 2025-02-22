package model

type User struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Age        int64  `json:"age"`
	Email      string `json:"email"`
	Contact    string `json:"contact"`
	Manager    bool   `json:"manager"`
	Department string `json:"department"`
}

type Login_system struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
