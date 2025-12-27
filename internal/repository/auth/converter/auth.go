package converter

import (
	"github.com/AndreiShkolnyi/go-auth/internal/repository/auth/model"
	"github.com/AndreiShkolnyi/go-auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromRepo(user *model.Auth) (*auth_v1.GetResponse, error) {

	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	} else {
		updatedAt = nil
	}
	return &auth_v1.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}, nil
}
