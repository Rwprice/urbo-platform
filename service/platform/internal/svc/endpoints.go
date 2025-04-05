package svc

import (
	"github.com/gin-gonic/gin"
)

type SVC struct {
	Manager *Manager
}

func (s *SVC) Register(engine *gin.Engine) {
	engine.GET("/ping", s.Ping)
}

func (s *SVC) Ping(c *gin.Context) {
	r := s.Manager.Pong()
	c.JSON(200, gin.H{
		"message": r,
	})
}
