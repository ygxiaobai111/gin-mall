package dao

import (
	"context"
	"gin-mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) CreateAddress(address *model.Address) error {
	return dao.DB.Model(model.Address{}).Create(&address).Error
}

func (dao *AddressDao) GetAddressByAid(aId, uId uint) (address *model.Address, err error) {
	err = dao.DB.Model(model.Address{}).Where("id=? AND user_id=?", aId, uId).First(&address).Error
	return
}

func (dao *AddressDao) UpdateByAddress(address *model.Address) (err error) {
	err = dao.DB.Model(model.Address{}).Save(&address).Error
	return
}

func (dao *AddressDao) CountAddressByUId(uId uint) (count int64, err error) {
	err = dao.DB.Model(model.Address{}).Where("user_id=?", uId).Count(&count).Error
	return
}
func (dao *AddressDao) GetAddressesByUId(uId uint) (addresses []*model.Address, err error) {
	err = dao.DB.Model(model.Address{}).Where("user_id=?", uId).Find(&addresses).Error
	return
}
func (dao *AddressDao) DelAddressByUId(address *model.Address, uId uint) (err error) {
	return dao.DB.Model(model.Address{}).Where("user_id=?", uId).Delete(&address).Error

}
