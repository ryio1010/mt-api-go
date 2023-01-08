package ports

import (
	"context"
	"mt-api-go/usecase/model"
)

type MusclePartInputPort interface {
	GetAllMusclePart(ctx context.Context) error
}

type MusclePartOutputPort interface {
	OutputMuscleParts(muscleParts []*model.MusclePart) error
	OutputError(err error) error
}
