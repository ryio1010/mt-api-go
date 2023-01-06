package interactors

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	models "mt-api-go/domain/model"
	"mt-api-go/domain/service"
	"mt-api-go/usecase/model"
	"mt-api-go/usecase/ports"
)

type IUserUseCase interface {
	LoginUser(ctx context.Context, user *model.User) error
	InsertNewUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
}

type UserUseCase struct {
	op  ports.UserOutputPort
	svc service.IUserService
}

func NewUserUseCase(uop ports.UserOutputPort, us service.IUserService) ports.UserInputPort {
	return &UserUseCase{
		op:  uop,
		svc: us,
	}
}

func (u *UserUseCase) LoginUser(ctx context.Context, user *model.User) error {
	// ログインユーザーの取得
	ms, err := u.svc.FindUserById(ctx, string(user.ID))

	if err != nil {
		return u.op.OutputError(err)
	}

	// パスワード認証
	err = bcrypt.CompareHashAndPassword([]byte(ms.Password), []byte(user.Password))

	if err != nil {
		return u.op.OutputError(err)
	}

	return u.op.OutputUser(model.UserFromDomainModel(ms))
}

func (u *UserUseCase) InsertNewUser(ctx context.Context, user *model.User) error {
	// パスワードハッシュ化
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	insertionTarget := models.MUser{
		Userid:   string(user.ID),
		Username: user.Name,
		Password: string(passwordHash),
	}

	err = u.svc.InsertNewUser(ctx, &insertionTarget)

	if err != nil {
		return u.op.OutputError(err)
	}

	return u.op.OutputUser(model.UserFromDomainModel(&insertionTarget))
}

func (u *UserUseCase) UpdateUser(ctx context.Context, user *model.User) error {
	updateTarget := models.MUser{
		Userid:   string(user.ID),
		Username: user.Name,
		Password: user.Password,
	}
	err := u.svc.UpdateUser(ctx, &updateTarget)

	if err != nil {
		return u.op.OutputError(err)
	}

	return u.op.OutputUser(model.UserFromDomainModel(&updateTarget))
}
