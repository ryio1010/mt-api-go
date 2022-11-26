package http

import "mt-api-go/usecase"

type userHandler struct {
	usecase usecase.IUserUseCase
}

func NewUserHandler(uu usecase.IUserUseCase) *userHandler {
	return &userHandler{
		usecase: uu,
	}
}

func (uh *userHandler) FindUserById() {

}
