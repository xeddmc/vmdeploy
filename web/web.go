package web

import (
	"github.com/gin-gonic/gin"
	"github.com/niemal/uman"
)

type Web struct {
	um     *uman.UserManager
	port   string
	engine *gin.Engine
}

func New() *Web {
	w := &Web{
		um:     uman.New("users.db"),
		port:   "8080",
		engine: gin.Default(),
	}

	w.engine.LoadHTMLGlob("web/templates/*")
	w.um.Register("admin", "default")

	return w
}

func (w *Web) Run() {
	w.Home()
	w.LogHandler()

	w.engine.Run(":" + w.port)
}