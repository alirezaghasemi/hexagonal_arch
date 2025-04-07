package general

import (
	"encoding/json"
	"fmt"

	"github.com/alirezaghasemi/hexagonal_arch/internal/entity"
	"github.com/alirezaghasemi/hexagonal_arch/internal/repo/port"
)

type userUseCase struct {
	UserRepo port.UserRepository
	Cache    port.CacheRepository
}

func NewUserUseCase(repo port.UserRepository, cache port.CacheRepository) UserUseCase {
	return &userUseCase{UserRepo: repo, Cache: cache}
}

func (uc *userUseCase) CreateUser(user entity.User) (entity.User, error) {
	user, err := uc.UserRepo.Create(user)
	if err != nil {
		return entity.User{}, err
	}

	//cache user
	userJSON, _ := json.Marshal(user)
	uc.Cache.Set(fmt.Sprintf("user:%s", user.Email), string(userJSON))

	return user, nil
}

func (uc *userUseCase) GetUserByEmail(email string) (entity.User, error) {
	// check cache
	cacheData, err := uc.Cache.Get(fmt.Sprintf("user:%s", email))
	if err == nil && cacheData != "" {
		var u entity.User
		err = json.Unmarshal([]byte(cacheData), &u)
		if err == nil {
			return u, nil
		}
	}

	// fallback to db
	user, err := uc.UserRepo.FindByEmail(email)
	if err != nil {
		return entity.User{}, err
	}

	// cache user
	userJSON, _ := json.Marshal(user)
	uc.Cache.Set(fmt.Sprintf("user:%s", user.Email), string(userJSON))

	return user, nil
}
