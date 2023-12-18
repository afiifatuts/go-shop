package app

import "github.com/afiifatuts/go-shop/app/models"

//disini nantinya method untuk meregister models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	//meregister model apa saja yang mau di migration ke method initializeDb yang ada di server.go
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
	}
}
