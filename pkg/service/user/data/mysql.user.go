package data

import (
	"database/sql"
	"github.com/alisyahbana/efishery-test/pkg/common/database"
	"github.com/jmoiron/sqlx"
)

type MysqlUserData struct {
}

const (
	CreateUserQuery = `INSERT INTO user (username, password, phone) VALUES (?, ?, ?)`
	GetUserQuery    = `SELECT username, password, phone, created_at, updated_at FROM user WHERE username = ? LIMIT 1`
	//UpdateUserQuery = `UPDATE user SET email = ?, address= ? WHERE username = ?`
	//DeleteUserQuery = `DELETE FROM user WHERE username = ?`
)

type MysqlUserStatement struct {
	CreateUserQuery *sqlx.Stmt
	GetUserQuery    *sqlx.Stmt
	//UpdateUserQuery *sqlx.Stmt
	//DeleteUserQuery *sqlx.Stmt
}

var stmt MysqlUserStatement

func init() {
	stmt.CreateUserQuery = database.Prepare(database.GetDBMaster(), CreateUserQuery)
	stmt.GetUserQuery = database.Prepare(database.GetDBMaster(), GetUserQuery)
	//stmt.UpdateUserQuery = database.Prepare(database.GetDBMaster(), UpdateUserQuery)
	//stmt.DeleteUserQuery = database.Prepare(database.GetDBMaster(), DeleteUserQuery)
}

func (m MysqlUserData) CreateUser(user User) (uint64, error) {
	newUser, err := stmt.CreateUserQuery.Exec(
		user.Username,
		user.Password,
		user.Phone,
	)
	if err != nil {
		return 0, err
	}

	newUserId, _ := newUser.LastInsertId()

	return uint64(newUserId), nil
}

func (m MysqlUserData) GetUser(username string) (*User, error) {
	var user User
	err := stmt.GetUserQuery.Get(
		&user,
		username,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

//func (m MysqlUserData) UpdateUser(payload UpdatePayload) error {
//	_, err := stmt.UpdateUserQuery.Exec(
//		payload.Email,
//		payload.Address,
//		payload.Username,
//	)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//func (m MysqlUserData) DeleteUser(username string) error {
//	_, err := stmt.DeleteUserQuery.Exec(
//		username,
//	)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
