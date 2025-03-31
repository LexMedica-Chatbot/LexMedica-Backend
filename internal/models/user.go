package models

import "time"

type User struct {
	ID        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Verified  bool      `db:"verified"`
	CreatedAt time.Time `db:"created_at"`
}
