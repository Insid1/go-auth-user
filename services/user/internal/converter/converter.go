package converter

import (
	"github.com/Insid1/with-auth/pkg/grpc/user_v1"

	"github.com/Insid1/with-auth/user/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromModel(user *model.User) *user_v1.User {
	return &user_v1.User{
		Id:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ToModelFromUser(user *user_v1.User) *model.User {
	return &model.User{
		ID:        user.GetId(),
		Email:     user.GetEmail(),
		Username:  user.GetUsername(),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
	}
}
