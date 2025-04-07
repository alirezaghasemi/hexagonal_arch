package general

import "github.com/alirezaghasemi/hexagonal_arch/internal/entity"

type UserUseCase interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}
