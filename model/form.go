package model

import "time"

type Form struct {
	Id        int64
	Name      string
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
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
func (formService *Form) DeleteItem(form *Form) error {
	return DB.Where("id =?", formService.Id).Delete(form).Error
}

func GetList(list *[]Form) error {
	return DB.Order("id asc").Find(list).Error
}
