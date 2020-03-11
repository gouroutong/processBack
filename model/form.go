package model

type Form struct {
	Id      int64
	Content string `gorm:"type:text"`
}

func (formService *Form) NewOrUpdate(form *Form) error {
	if formService.Id == 0 {
		return DB.Create(&formService).Error
	} else {
		return DB.Model(formService).Where("id =?", formService.Id).Update(&formService).Error
	}
}
func (formService *Form) GetItem(form *Form) error {
	return DB.Where("id =?", formService.Id).Find(form).Error
}

func GetList(list *[]Form) error {
	return DB.Order("id asc").Find(list).Error
}
