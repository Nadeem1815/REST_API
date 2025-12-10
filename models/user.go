package models

import (
	"errors"

	"github.com/Nadeem1815/rest-api/db"
	"github.com/Nadeem1815/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email,password)
			 VALUES (?,?)
			 `
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT id, password 
	FROM users 
	WHERE email =?
	`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	if err := row.Scan(&u.ID, &retrievedPassword); err != nil {
		return errors.New("Credentials invalid")
	}

	if passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword); !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
