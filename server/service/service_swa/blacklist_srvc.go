package service_swa

import "context"

type TBlacklist interface {
	CreateBlacklist(_ context.Context, jwtList JwtBlacklist) (err error)
}


func (imp *impl) CreateBlacklist(_ context.Context, jwtList JwtBlacklist) (err error) {
	err = imp.gormDB.Create(&jwtList).Error
	if err != nil {
		return
	}
	return err
}
