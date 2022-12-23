package usecase

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	models "mt-api-go/domain/model"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
)

type IUserUseCase interface {
	LoginUser(ctx context.Context, user *model.User) (*model.User, error)
	InsertNewUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
}

type userUseCase struct {
	svc service.IUserService
}

func NewUserUseCase(us service.IUserService) IUserUseCase {
	return &userUseCase{
		svc: us,
	}
}

func (uu *userUseCase) LoginUser(ctx context.Context, user *model.User) (*model.User, error) {
	// ログインユーザーの取得
	ms, err := uu.svc.FindUserById(ctx, string(user.ID))

	if err != nil {
		return nil, err
	}

	// パスワード認証
	err = bcrypt.CompareHashAndPassword([]byte(ms.Password), []byte(user.Password))

	if err != nil {
		return nil, err
	}

	return model.UserFromDomainModel(ms), nil
}

func (uu *userUseCase) InsertNewUser(ctx context.Context, user *model.User) error {
	// パスワードハッシュ化
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	insertionTarget := models.MUser{
		Userid:   string(user.ID),
		Username: user.Name,
		Password: string(passwordHash),
	}
	err = uu.svc.InsertNewUser(ctx, &insertionTarget)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUseCase) UpdateUser(ctx context.Context, user *model.User) error {
	updateTarget := models.MUser{
		Userid:   string(user.ID),
		Username: user.Name,
		Password: user.Password,
	}
	err := uu.svc.UpdateUser(ctx, &updateTarget)

	if err != nil {
		return err
	}

	return nil
}
