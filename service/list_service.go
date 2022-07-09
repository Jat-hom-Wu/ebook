package service

import "github.com/Jat-hom-Wu/ebook/model"

func AddList(tittle string, name string) error{
	return model.ListAdd(name, tittle, false)
}

func DeleteList(id string) error {
	return model.ListDelete(id)
}

func UpdateList(status bool, id int) error{
	return model.ListUpdate(status, id)
}

func GetList(name string) ([]model.ListTable, error) {
	return model.ListCheck(name)
}

