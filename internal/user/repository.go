package user

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]*User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id int, user *User) error
	Delete(ctx context.Context, id int) error
}

type UserREPO struct {
	db *gorm.DB
}

func NewUserREPO(db *gorm.DB) *UserREPO {
	return &UserREPO{
		db: db,
	}
}

func (r *UserREPO) GetUsers(ctx context.Context) ([]*User, error) {
	var users []User
	err := r.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	userPtrs := make([]*User, 0, len(users))
	for i := range users {
		userPtrs = append(userPtrs, &users[i])
	}
	return userPtrs, nil
}

func (r *UserREPO) GetByID(ctx context.Context, id int) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserREPO) Create(ctx context.Context, user *User) error {

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserREPO) Update(ctx context.Context, id int, user *User) error {
	user.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"email": user.Email,
			"pass":  user.Pass,
		}).Error
}

func (r *UserREPO) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&User{}, id).Error
}
