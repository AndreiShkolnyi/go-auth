package repository

import (
	"context"

	"github.com/AndreiShkolnyi/go-auth/pkg/auth_v1"
)

type AuthRepository interface {
	Create(ctx context.Context, user *auth_v1.CreateRequest) (int64, error)
	Get(ctx context.Context, id int64) (*auth_v1.GetResponse, error)
	Update(ctx context.Context, user *auth_v1.UpdateRequest) (int64, error)
	Delete(ctx context.Context, id int64) (int64, error)
}
