package rdb

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

type BodyCompRepository struct {
	DB *sql.DB
}

func NewBodyCompRepository(db *sql.DB) repository.IBodyCompRepository {
	return &BodyCompRepository{
		DB: db,
	}
}

func (b *BodyCompRepository) SelectBodyCompById(ctx context.Context, userId string, trainingDate string) (*model.TBodycomp, error) {
	query := fmt.Sprintf("%s = ? AND %s = ?", model.TBodycompColumns.Userid, model.TTraininglogColumns.Trainingdate)
	orderBy := fmt.Sprintf("%s ?", model.TBodycompColumns.Upddate)
	bodyComp, err := model.TBodycomps(
		qm.Where(query, userId, trainingDate),
		qm.OrderBy(orderBy, "DESC"),
	).One(ctx, b.DB)

	if err != nil {
		log.Error().Msg(err.Error())
	}

	fmt.Print("selectBodyCompInfo")
	fmt.Println(bodyComp)
	return bodyComp, err
}

func (b *BodyCompRepository) SelectLatestBodyComp(ctx context.Context, userId string) (*model.TBodycomp, error) {
	query := fmt.Sprintf("%s = ?", model.TBodycompColumns.Userid)
	bodyComp, err := model.TBodycomps(
		qm.Where(query, userId),
	).One(ctx, b.DB)

	if err != nil {
		log.Error().Msg(err.Error())
	}

	fmt.Print("selectBodyCompInfo")
	fmt.Println(bodyComp)
	return bodyComp, err
}

func (b *BodyCompRepository) InsertBodyComp(ctx context.Context, target *model.TBodycomp) (*model.TBodycomp, error) {
	target.Regid = null.StringFrom(target.Userid)
	target.Regdate = null.TimeFrom(time.Now())
	target.Updid = null.StringFrom(target.Userid)
	target.Upddate = null.TimeFrom(time.Now())
	target.Version = 1

	err := target.Insert(ctx, b.DB, boil.Infer())
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return target, err
}

func (b *BodyCompRepository) UpdateBodyComp(ctx context.Context, target *model.TBodycomp) (*model.TBodycomp, error) {
	target.Updid = null.StringFrom(target.Userid)
	target.Upddate = null.TimeFrom(time.Now())
	target.Version = target.Version + 1

	_, err := target.Update(ctx, b.DB, boil.Infer())
	if err != nil {
		log.Error().Msg(err.Error())
		return target, err
	}
	return target, err
}
