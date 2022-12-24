package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
)

type TrainingLogRepository struct {
	DB *sql.DB
}

func NewTrainingLogRepository(db *sql.DB) repository.ITrainingLogRepository {
	return &TrainingLogRepository{
		DB: db,
	}
}

func (lr *TrainingLogRepository) SelectTrainingLogById(ctx context.Context, userId string) ([]*model.TTraininglog, error) {
	id := fmt.Sprintf("%s = ?", model.MMenuColumns.Userid)
	trainingLogList, err := model.TTraininglogs(
		qm.Where(id, userId),
	).All(ctx, lr.DB)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Print("selectLogInfo")
	fmt.Println(trainingLogList)
	return trainingLogList, err
}

func (lr *TrainingLogRepository) InsertTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error) {
	return nil, nil
}

func (lr *TrainingLogRepository) UpdateTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error) {
	return nil, nil
}

func (lr *TrainingLogRepository) DeleteTrainingLog(ctx context.Context, logId int) (int, error) {
	return 0, nil
}

func (lr *TrainingLogRepository) DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) (int, string, error) {
	return 0, "", nil
}
