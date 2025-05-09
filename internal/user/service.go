package user

import "context"

type Service interface {
	GetUsers(ctx context.Context) ([]*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id int, user *User) error
	DeleteUser(ctx context.Context, id int) error
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers(ctx context.Context) ([]*User, error) {
	return s.repo.GetUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) CreateUser(ctx context.Context, user *User) error {
	return s.repo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user *User) error {
	return s.repo.Update(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
