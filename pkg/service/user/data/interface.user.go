package data

import "time"

type User struct {
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Phone     string    `json:"phone" db:"phone"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type UserClaims struct {
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Timestamp string `json:"timestamp"`
}

type UpdatePayload struct {
	Email    string `json:"email"`
	Address  string `json:"address"`
	Username string `json:"username"`
}

type UserData interface {
	CreateUser(user User) (uint64, error)
	GetUser(username string) (*User, error)
	GetUserByPhoneAndPassword(phone string, password string) (*User, error)
	//UpdateUser(payload UpdatePayload) error
	//DeleteUser(username string) error
}
