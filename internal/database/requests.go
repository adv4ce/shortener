package database

import (
	"context"
	"shortener/internal/services"
	"time"

	"gorm.io/gorm"
)

type DBRepo struct {
	db *gorm.DB
}

func InitDBRepo(db *gorm.DB) *DBRepo {
	return &DBRepo{
		db: db,
	}
}

func (r *DBRepo) GetCode(ctx context.Context, url string) (string, error) {
	furl, err := gorm.G[Url](r.db).Where("url = ?", url).First(ctx)

	if err != nil {
		newUrl := Url{
			URL: url,
		}
		_ = gorm.G[Url](r.db).Create(ctx, &newUrl)

		code := services.CreateShortCode(newUrl.ID)

		_, _ = gorm.G[Url](r.db).Where("url = ?", url).Update(ctx, "code", code)
		return code, nil
	}

	return furl.Code, nil
}

func (r *DBRepo) GetURL(ctx context.Context, code string) (*Url, error) {
	var err error
	var u Url

	if u, err = gorm.G[Url](r.db).Where("code = ?", code).First(ctx); err != nil {
		return nil, err
	}

	u.Clicks++
	u.LastClick = time.Now()

	_, err = gorm.G[Url](r.db).Where("code = ?", code).Updates(ctx, u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
