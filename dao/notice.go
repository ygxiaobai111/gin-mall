package dao

import (
	"context"
	"fmt"
	"gin-mall/model"
	"gorm.io/gorm"
)

// NewNoticeDao

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNewNoticeDaoByDB 根据id获取Notice
func (dao *NoticeDao) GetNewNoticeDaoByDB(id uint) (notice *model.Notice, err error) {
	fmt.Println(id)
	err = dao.DB.Model(&model.Notice{}).Where("id = ?", id).First(&notice).Error
	return
}
