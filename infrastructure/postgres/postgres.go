package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"mt-api-go/domain/model"
	"mt-api-go/domain/repository"
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
	return model.MUsers(
		qm.Where(id, userId),
	).One(ctx, ur.DB)
}
