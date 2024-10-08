package types

import "database/sql"

type Student struct {
	Id       int64  `json:"id"`
	Name     string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `gorm:"size:100" json:"password" validate:"required"`
}

type Sqlite struct {
	Db *sql.DB
}

type Mysql struct {
	Db *sql.DB
}
