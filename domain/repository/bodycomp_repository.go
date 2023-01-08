package repository

import (
	"context"
	"mt-api-go/domain/model"
)

type IBodyCompRepository interface {
	SelectBodyCompById(ctx context.Context, userId string, trainingDate string) (*model.TBodycomp, error)
	SelectLatestBodyComp(ctx context.Context, userId string) (*model.TBodycomp, error)
	InsertBodyComp(ctx context.Context, model *model.TBodycomp) (*model.TBodycomp, error)
	UpdateBodyComp(ctx context.Context, model *model.TBodycomp) (*model.TBodycomp, error)
}
