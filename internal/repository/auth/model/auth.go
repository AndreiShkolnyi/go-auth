package model

import (
	"database/sql"
	"time"

	"github.com/AndreiShkolnyi/go-auth/pkg/auth_v1"
)

type Auth struct {
	ID              int64        `db:"id"`
	Name            string       `db:"name"`
	Email           string       `db:"email"`
	Password        string       `db:"password"`
	PasswordConfirm string       `db:"password_confirm"`
	Role            auth_v1.Role `db:"role"`
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       sql.NullTime `db:"updated_at"`
}
