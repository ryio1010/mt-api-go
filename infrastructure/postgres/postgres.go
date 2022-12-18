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

type userRepository struct {
	DB *sql.DB
}

func NewRoomRepository(db *sql.DB) repository.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) SelectUserById(ctx context.Context, userId string) (*model.MUser, error) {
	id := fmt.Sprintf("%s = ?", model.MUserColumns.Userid)
	user, err := model.MUsers(
		qm.Where(id, userId),
	).One(ctx, ur.DB)
	if err != nil {
		log.Error().Msg(err.Error())
	}
	fmt.Print("selectUserInfo")
	fmt.Println(user)
	return user, err
}

func (ur *userRepository) InsertNewUser(ctx context.Context, user *model.MUser) error {
	user.Regid = null.StringFrom(user.Userid)
	user.Regdate = null.TimeFrom(time.Now())
	user.Updid = null.StringFrom(user.Userid)
	user.Upddate = null.TimeFrom(time.Now())
	user.Version = 1

	err := user.Insert(ctx, ur.DB, boil.Infer())
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *model.MUser) error {
	user.Regid = null.StringFrom(user.Userid)
	user.Regdate = null.TimeFrom(time.Now())
	user.Updid = null.StringFrom(user.Userid)
	user.Upddate = null.TimeFrom(time.Now())
	user.Version = 1

	_, err := user.Update(ctx, ur.DB, boil.Infer())
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}
