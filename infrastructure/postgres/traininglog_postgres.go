package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
	"time"
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
	target.Regid = null.StringFrom(target.Userid)
	target.Regdate = null.TimeFrom(time.Now())
	target.Updid = null.StringFrom(target.Userid)
	target.Upddate = null.TimeFrom(time.Now())
	target.Version = 1

	err := target.Insert(ctx, lr.DB, boil.Infer())
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return target, nil
}

func (lr *TrainingLogRepository) UpdateTrainingLog(ctx context.Context, target *model.TTraininglog) (*model.TTraininglog, error) {
	target.Updid = null.StringFrom(target.Userid)
	target.Upddate = null.TimeFrom(time.Now())
	target.Version = target.Version + 1

	_, err := target.Update(ctx, lr.DB, boil.Infer())
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	return target, nil
}

func (lr *TrainingLogRepository) DeleteTrainingLog(ctx context.Context, id int) (int, error) {
	logId := fmt.Sprintf("%s = ?", model.TTraininglogColumns.Logid)
	deleteTarget, err := model.TTraininglogs(
		qm.Where(logId, id),
	).One(ctx, lr.DB)

	if err != nil {
		log.Error().Msg(err.Error())
		return 0, err
	}

	_, err = deleteTarget.Delete(ctx, lr.DB)
	if err != nil {
		log.Error().Msg(err.Error())
		return 0, err
	}
	return id, nil
}

func (lr *TrainingLogRepository) DeleteTrainingLogByMenuIdAndDate(ctx context.Context, menuId int, trainingDate string) (int, string, error) {
	query := fmt.Sprintf("%s = ? AND %s = ?", model.TTraininglogColumns.Menuid, model.TTraininglogColumns.Trainingdate)
	deleteTargetList, err := model.TTraininglogs(
		qm.Where(query, menuId, trainingDate),
	).All(ctx, lr.DB)

	if err != nil {
		log.Error().Msg(err.Error())
		return 0, "", err
	}

	for _, target := range deleteTargetList {
		_, err = target.Delete(ctx, lr.DB)

		if err != nil {
			log.Error().Msg(err.Error())
			return 0, "", err
		}
	}

	return menuId, trainingDate, nil
}
