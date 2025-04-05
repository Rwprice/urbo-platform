package main

import (
	"github.com/gin-gonic/gin"
	"urbo-platform/service/platform/internal/svc"
)

func main() {
	r := gin.Default()
	manager := svc.NewManager(r)

	s := svc.SVC{
		Manager: manager,
	}
	s.Register(r)

	err := r.Run()
	if err != nil {
		return
	}
}
