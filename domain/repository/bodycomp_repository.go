package repository

import (
	"context"
)

type IBodyCompRepository interface {
	selectBodyCompById(ctx context.Context)
	selectLatestBodyComp(ctx context.Context)
	insertBodyComp(ctx context.Context)
	updateBodyComp(ctx context.Context)
}
