package model

import (
	"log"
)

type Repository interface {
	GetAll() []interface{}
	GetById(interface{}) interface{}
	Save(interface{})
}

type UserRepository struct {
	con *connection
}

func GetUserRepository() *UserRepository {
	return &UserRepository{con: DatabaseConnect()}
}

func (userRepo *UserRepository) GetAll() ([]*UserData, error) {
	rows, err := userRepo.con.Query("SELECT * FROM gouser")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*UserData, 0)
	for rows.Next() {
		user := new(UserData)
		err := rows.Scan(&user.UserName, &user.UserPassword)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	log.Printf("%+v \n", users)
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepo *UserRepository) GetById(id int) (*UserData, error) {
	row, err := userRepo.con.Query("SELECT * FROM gouser WHERE gouser.id=$1", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	user := new(UserData)
	err = row.Scan(&user.UserName, &user.UserPassword)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *UserRepository) Save(user *UserData) error {
	row, err := userRepo.con.Query("INSERT INTO gouser (username, user_password) VALUES ($1, $2 )", user.UserName, user.UserPassword)
	if err != nil {
		return err
	}
	log.Printf("insert new user %+v \n", row)
	defer row.Close()
	return nil
}
