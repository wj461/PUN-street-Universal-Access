package domain

import (
	"context"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

type UserRepo interface {
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error)
	Login(ctx context.Context, email string, password string) (string, int64, error)
	RegisterUser(ctx context.Context, user *swagger.RegisterInfo, authority string) (int, error)
	RegisterStore(ctx context.Context, storeInfo swagger.StoreRegisterInfo, id int) error
	CheckEmail(ctx context.Context, email string) (bool, error)
	IsUserBanned(ctx context.Context, email string) (bool, error)
}

type UserUsecase interface {
	GetByID(ctx context.Context, id int64) (*swagger.UserData, error)
	GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error)
	Login(ctx context.Context, email string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) error
	RegisterUser(ctx context.Context, user *swagger.RegisterInfo) error
	CheckEmail(ctx context.Context, email string) (bool, error)
	GetUserIdByCookie(ctx context.Context, token string) (int64, error)
}
