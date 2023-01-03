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

type TrainingMenuRepository struct {
	DB *sql.DB
}

func NewTrainingMenuRepository(db *sql.DB) repository.ITrainingMenuRepository {
	return &TrainingMenuRepository{
		DB: db,
	}
}

func (mr *TrainingMenuRepository) SelectMenuByUserId(ctx context.Context, userId string) ([]*model.MMenu, error) {
	id := fmt.Sprintf("%s = ?", model.MMenuColumns.Userid)
	trainingMenuList, err := model.MMenus(
		qm.Where(id, userId),
	).All(ctx, mr.DB)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Print("selectMenuInfo")
	fmt.Println(trainingMenuList)
	return trainingMenuList, err
}

func (mr *TrainingMenuRepository) InsertMenu(ctx context.Context, menuEntity *model.MMenu) (*model.MMenu, error) {
	menuEntity.Regid = menuEntity.Userid
	menuEntity.Regdate = null.TimeFrom(time.Now())
	menuEntity.Updid = menuEntity.Userid
	menuEntity.Upddate = null.TimeFrom(time.Now())
	menuEntity.Version = 1

	err := menuEntity.Insert(ctx, mr.DB, boil.Infer())
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return menuEntity, err
}

func (mr *TrainingMenuRepository) DeleteMenu(ctx context.Context, id int) (int, error) {
	menuId := fmt.Sprintf("%s = ?", model.MMenuColumns.Menuid)
	trainingMenu, err := model.MMenus(
		qm.Where(menuId, id),
	).One(ctx, mr.DB)

	_, err = trainingMenu.Delete(ctx, mr.DB)
	if err != nil {
		log.Error().Msg(err.Error())
		return 0, err
	}
	return id, err
}
