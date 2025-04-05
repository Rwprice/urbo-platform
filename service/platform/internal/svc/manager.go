package svc

import (
	"github.com/gin-gonic/gin"
)

type Manager struct {
}

func NewManager(engine *gin.Engine) *Manager {
	return &Manager{}
}

func (m *Manager) Pong() string {
	return "pong"
}
