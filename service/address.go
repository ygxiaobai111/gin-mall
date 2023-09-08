package service

import (
	"context"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{

		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		util.LogrusObj.Info("err: ", err) //保存日志
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Error:  err.Error(),
	}
}
func (service *AddressService) Update(ctx context.Context, aId string, uId uint) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(uint(addressId), uId)

	if err != nil {
		util.LogrusObj.Info("err: ", err) //保存日志
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if address.Name != service.Name {
		address.Name = service.Name
	}
	if address.Address != service.Address {
		address.Address = service.Address
	}
	if address.Phone != service.Phone {
		address.Phone = service.Phone
	}
	err = addressDao.UpdateByAddress(address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}

}
func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	var addresses []*model.Address
	var err error
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	count, err := addressDao.CountAddressByUId(uId)
	if err != nil {
		util.LogrusObj.Info("err: ", err) //保存日志
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	addresses, err = addressDao.GetAddressesByUId(uId)
	if err != nil {
		util.LogrusObj.Info("err: ", err) //保存日志
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildAddresses(addresses), uint(count))
}
func (service *AddressService) Delete(ctx context.Context, aId string, uId uint) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	var address *model.Address
	code := e.Success
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(uint(addressId), uId)
	if address == nil {
		code = e.ErrorAddressNotExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = addressDao.DelAddressByUId(address, uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
func (service *AddressService) Show(ctx context.Context, aId string, uId uint) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	code := e.Success

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(uint(addressId), uId)
	if err != nil {
		util.LogrusObj.Info("err: ", err) //保存日志
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}

}
