package entity

type user struct {
	id       string `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
	image    string `json:"image"`
}
