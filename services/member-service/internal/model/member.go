package model

type Member struct {
	Username string `json:"username"`
	Mid      int64  `json:"mid"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
