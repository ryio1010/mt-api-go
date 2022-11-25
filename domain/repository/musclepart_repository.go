package repository

import (
	"context"
)

type IMusclePartRepository interface {
	selectAllMusclePart(ctx context.Context)
}
