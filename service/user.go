package service

import (
	"context"
	"gin-mall/dao"
	"gin-mall/model"
	"gin-mall/pkg/e"
	"gin-mall/serializer"
	"gin-mall/util"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //前端验证
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}

	//将明文信息加密
	util.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}

	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{

		UserName: service.UserName,
		NickName: service.NickName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesDecoding("10000"), //金额初始化(加密)
	}
	//密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}

	}
	//用户创建
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error

	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}

}

func (service *UserService) Login(ctx context.Context) serializer.Response {

}
