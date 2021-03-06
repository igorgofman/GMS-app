package repository

import (
	"fmt"

	"github.com/igorgofman/gms-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user gym.SysUser) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password_hash) values ($1, $2, $3) RETURNING id", sysUsersTable)

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (gym.SysUser, error) {
	var user gym.SysUser
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", sysUsersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
