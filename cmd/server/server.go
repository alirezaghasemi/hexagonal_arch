package server

import (
	"github.com/alirezaghasemi/hexagonal_arch/config"
	"github.com/alirezaghasemi/hexagonal_arch/internal/container"
	"github.com/alirezaghasemi/hexagonal_arch/internal/server"
	"github.com/gin-gonic/gin"
)

func StartServe() {
	cfg := config.LoadConfig()
	c := container.NewContainer(cfg)
	r := gin.Default()
	server.RegisterRoutes(r, c.UserUC)

	r.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}
