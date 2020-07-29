package domain

type Login struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
