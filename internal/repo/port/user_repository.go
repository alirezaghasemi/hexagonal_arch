package port

import "github.com/alirezaghasemi/hexagonal_arch/internal/entity"

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
