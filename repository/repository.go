package repository

import "encoding/json"

type Repository interface {
	GetAll() []interface{}
	GetById(interface{}) interface{}
	Save(interface{})
	SaveAll([]interface{})
}

type PostData struct {
	userName     string
	userPassword string
}

func (data *PostData) ToStruct(rawData []byte) error {
	return json.Unmarshal(rawData, &data)
}
