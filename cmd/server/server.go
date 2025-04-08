package server

import (
	"log"

	"github.com/alirezaghasemi/hexagonal_arch/config"
	"github.com/alirezaghasemi/hexagonal_arch/internal/container"
	"github.com/alirezaghasemi/hexagonal_arch/internal/server"
	"github.com/gin-gonic/gin"
)

func StartServe() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	c := container.NewContainer(cfg)
	r := gin.Default()
	server.RegisterRoutes(r, c.UserUC)

	r.Run(cfg.Server.Host + ":" + cfg.Server.Port)
}
