package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
)

type MusclePartRepository struct {
	DB *sql.DB
}

func NewMusclePartRepository(db *sql.DB) repository.IMusclePartRepository {
	return &MusclePartRepository{
		DB: db,
	}
}

func (mr *MusclePartRepository) SelectAllMusclePart(ctx context.Context) ([]*model.MMusclepart, error) {
	muscleParts, err := model.MMuscleparts().All(ctx, mr.DB)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Println(muscleParts)
	return muscleParts, err
}
