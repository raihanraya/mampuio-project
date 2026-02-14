package models

import (
	"context"
	"errors"
	"mampuio-project/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name    string
	Balance float64
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (um *UserRepository) GetUserByName(name string) (user User) {
	ctx := context.Background()
	user, err := gorm.G[User](config.DB).Where("name = ?", name).First(ctx)
	if err != nil {
		panic("failed get data")
	}

	return user
}

func (um *UserRepository) WithdrawBalance(user User, amount float64) error {

	if user.Balance < amount {
		return errors.New("insufficient balance")
	}

	user.Balance -= amount

	config.DB.Save(&user)

	return nil
}

func (um *UserRepository) GetBalance(name string) int {
	var user User

	config.DB.First(&user, "name = ?", name)

	return int(user.Balance)
}
