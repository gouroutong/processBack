package model

import (
	"time"
)

type Node struct {
	Id        int64
	Name      string
	FormId    int64
	Act       string
	UserId    int64
	Username  string
	FormName  string
	ProcessId int64 `gorm:"index"`
}

type Process struct {
	Id        int64
	Name      string
	NodeList  []Node
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (processService *Process) NewOrUpdateProcess(process *Process) error {
	if processService.Id == 0 {
		return DB.Create(processService).Error
	} else {
		err := DB.Where("process_id = ?", processService.Id).Delete(&Node{}).Error
		if err != nil {
			return err
		}
		return DB.Model(processService).Where("id =?", processService.Id).Update(processService).Error
	}
}
func (processService *Process) GetProcessItem(process *Process) error {
	err := DB.Where("id =?", processService.Id).Find(process).Error
	if err != nil {
		return err
	}
	return DB.Where("process_id =?", processService.Id).Find(&process.NodeList).Error
}
func GetProcessList(list *[]Process) error {
	return DB.Order("id asc").Find(list).Error
}
func (processService *Process) DeleteProcessItem(process *Process) error {
	err := DB.Where("process_id = ?", processService.Id).Delete(&Node{}).Error
	if err != nil {
		return err
	}
	return DB.Where("id =?", processService.Id).Delete(process).Error
}
