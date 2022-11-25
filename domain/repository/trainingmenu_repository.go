package repository

import (
	"context"
)

type ITrainingMenuRepository interface {
	selectAllMusclePart(ctx context.Context)
}
