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

type FavoriteService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

// Show 商品收藏夹
func (service *FavoriteService) Show(ctx context.Context, uId uint) serializer.Response {
	favoritesDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	favorites, total, err := favoritesDao.ListFavoriteByUserId(uId, service.PageSize, service.PageNum)
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(total))
}

// Create 创建收藏夹
func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	bossDao := dao.NewUserDaoByDB(userDao.DB)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(int(service.ProductId))
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favorite := &model.Favorite{
		UserId:    uId,
		User:      *user,
		ProductId: service.ProductId,
		Product:   *product,
		BossId:    service.BossId,
		Boss:      *boss,
	}
	favoriteDao = dao.NewFavoriteDaoByDB(favoriteDao.DB)
	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Delete 删除收藏夹
func (service *FavoriteService) Delete(ctx context.Context, fId string, uId uint) serializer.Response {
	code := e.Success

	favoriteDao := dao.NewFavoriteDao(ctx)
	id, _ := strconv.Atoi(fId)
	err := favoriteDao.DeleteFavoriteById(uint(id), uId)
	if err != nil {
		util.LogrusObj.Infoln(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
	}
}
