package serializer

import (
	"context"
	"gin-mall/conf"

	dao2 "gin-mall/dao"
	model2 "gin-mall/model"
)

type Favorite struct {
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreatedAt     int64  `json:"create_at"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossID        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

// BuildFavorite 序列化收藏夹
func BuildFavorite(item1 *model2.Favorite, item2 *model2.Product, item3 *model2.User) Favorite {
	return Favorite{
		UserID:        item1.UserId,
		ProductID:     item1.ProductId,
		CreatedAt:     item1.CreatedAt.Unix(),
		Name:          item2.Name,
		CategoryID:    item2.CategoryId,
		Title:         item2.Title,
		Info:          item2.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + item2.ImgPath,
		Price:         item2.Price,
		DiscountPrice: item2.DiscountPrice,
		BossID:        item3.ID,
		Num:           item2.Num,
		OnSale:        item2.OnSale,
	}
}

// BuildFavorites 收藏夹列表
func BuildFavorites(ctx context.Context, items []*model2.Favorite) (favorites []Favorite) {
	productDao := dao2.NewProductDao(ctx)
	bossDao := dao2.NewUserDao(ctx)

	for _, fav := range items {
		product, err := productDao.GetProductById(int(fav.ProductId))
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(fav.BossId)
		if err != nil {
			continue
		}
		favorite := BuildFavorite(fav, product, boss)
		favorites = append(favorites, favorite)
	}
	return favorites
}
