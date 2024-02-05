package model

const (
	RoleUser    = "user"
	RoleManager = "manager"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ChatID   int64  `json:"chat_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
