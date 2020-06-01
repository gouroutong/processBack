package model

import (
	"errors"
	"time"
	"xProcessBackend/ws"
)

type Apply struct {
	Id            int64
	Form          []ApplyForm
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Status        int
	ProcessId     int64
	StartUserId   int64
	StartUsername string
	ProcessName   string
	NextRoleId    int64
}

type ApplyForm struct {
	Id      int64
	ApplyId int64
	Content string `gorm:"type:text"`
}

func (applyService *Apply) NewApply(apply *Apply, userId int64) error {
	if (applyService.NextRoleId != 0) || (applyService.NextRoleId == userId) {
		ws.ConnIns.Update()
	}

	if applyService.Id == 0 {
		return DB.Create(applyService).Error
	} else {
		var (
			err error
		)
		err = DB.Model(applyService).Where("id =?", applyService.Id).First(&apply).Error
		if err != nil {
			return err
		}

		if apply.Status == 2 {
			return errors.New("已经被审核过了")
		}
		return DB.Model(applyService).Where("id =? AND status=?", applyService.Id, 0).Update(&applyService).Error
	}
}

func (applyService *Apply) GetApplyItem(apply *Apply) error {
	err := DB.Where("id =?", applyService.Id).Find(apply).Error
	if err != nil {
		return err
	}
	return DB.Where("apply_id =?", applyService.Id).Find(&apply.Form).Error
}

func GetApplyList(list *[]Apply, userId int64) error {
	return DB.Order("id asc").Where("start_user_id=?", userId).Find(list).Error
}
func GetAuditApplyList(list *[]Apply, userId int64) error {
	return DB.Order("id asc").Where("next_role_id=?", userId).Find(list).Error
}

func (applyService *Apply) DeleteApplyItem(apply *Apply) error {
	return DB.Where("id =?", applyService.Id).Delete(apply).Error
}
