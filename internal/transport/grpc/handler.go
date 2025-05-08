package grpc

import (
	userpb "github.com/Suhach/protoc-cont/proto"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}
