package container

import (
	"github.com/alirezaghasemi/hexagonal_arch/config"
	"github.com/alirezaghasemi/hexagonal_arch/internal/entity"
	"github.com/alirezaghasemi/hexagonal_arch/internal/provider"
	"github.com/alirezaghasemi/hexagonal_arch/internal/repo/adapter"
	"github.com/alirezaghasemi/hexagonal_arch/internal/usecase/general"
)

type Container struct {
	UserUC general.UserUseCase
}

func NewContainer(cfg *config.Config) *Container {

	db := provider.InitPostgres(cfg)
	db.AutoMigrate(&entity.User{})

	redis := provider.InitRedis(cfg)

	userRepo := adapter.NewGormUserRepo(db)
	userCache := adapter.NewRedisCache(redis)

	return &Container{
		UserUC: general.NewUserUseCase(userRepo, userCache),
	}
}
