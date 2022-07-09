package model



type ListTable struct{
	Name string `json:"name" gorm:"column:name"`
	Title string `json:"title" gorm:"column:title"`
	Status bool `json:"status" gorm:"column:status"`
	Id int `json:"id" gorm:"column:id"`
}

func ListAdd(name,title string, status bool) error{
	obj := ListTable{Name:name,Title:title,Status:status}
	return DB.Select("name", "title", "status").Create(&obj).Error
}

func ListDelete(id string) error{
	result := DB.Where("id = ?", id).Delete(&ListTable{})
	return result.Error
}

func ListUpdate(status bool, id int) error{
	result := DB.Model(&ListTable{}).Where("id = ?", id).Select("status").Updates(ListTable{Status: status})
	// result := DB.Model(&ListTable{}).Select("id = ?", id).Update(ListTable{Status: status})
	return result.Error
}

func ListCheck(name string) ([]ListTable, error){
	listtables := []ListTable{}
	result := DB.Where("name = ?", name).Find(&listtables)
	if result.Error != nil{
		return listtables,result.Error
	}
	return listtables,nil
}



