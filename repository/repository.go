package repository

type Repository interface {
	GetAll() []interface{}
	GetById(interface{}) interface{}
	Save(interface{})
	SaveAll([]interface{})
}
