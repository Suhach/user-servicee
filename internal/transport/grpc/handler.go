package grpc

import (
	"context"
	userpb "github.com/Suhach/protoc-cont/proto"
	"github.com/Suhach/user-servicee/internal/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetUsers(ctx context.Context, _ *emptypb.Empty) (*userpb.GetUsersResponse, error) {
	users, err := h.svc.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var res []*userpb.User
	for _, user := range users {
		res = append(res, &userpb.User{
			Id:    uint32(user.ID),
			Email: user.Email,
			Pass:  user.Pass,
		})
	}

	return &userpb.GetUsersResponse{User: res}, nil
}

func (h *Handler) GetUserByID(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := h.svc.GetUserByID(ctx, int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    uint32(user.ID),
			Email: user.Email,
			Pass:  user.Pass,
		},
	}, nil

}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user := &user.User{
		Email: req.Email,
		Pass:  req.Pass,
	}

	err := h.svc.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	res := &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    uint32(user.ID),
			Email: user.Email,
			Pass:  user.Pass,
		},
	}
	return res, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	user := &user.User{
		ID:        int(req.Id),
		Pass:      req.Pass,
		Email:     req.Email,
		UpdatedAt: time.Now(),
	}

	err := h.svc.UpdateUser(ctx, int(req.Id), user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user: %v", err)
	}
	res := &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    uint32(user.ID),
			Pass:  user.Pass,
			Email: user.Email,
			//UpdatedAt: user.UpdatedAt,
		},
	}
	return res, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.svc.DeleteUser(ctx, int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}
	return &userpb.DeleteUserResponse{
		Success: true,
	}, nil
}
