package repository

import (
	"context"
	"mt-api-go/domain/model"
)

type IMusclePartRepository interface {
	SelectAllMusclePart(ctx context.Context) ([]*model.MMusclepart, error)
}
