package model

import "encoding/json"

type UserData struct {
	Id           int64  `db:"id" json:"id"`
	UserName     string `db:"username" json:"username"`
	UserPassword string `db:"user_password" json:"password"`
}

func (data *UserData) ToJson() ([]byte, error) {
	return json.Marshal(data)
}

func (data *UserData) ToObject(rawData []byte) error {
	return json.Unmarshal(rawData, &data)
}
