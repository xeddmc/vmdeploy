package web

import (
	"github.com/gin-gonic/gin"
	"github.com/niemal/uman"
	"net/http"
)

func home(c *gin.Context, um *uman.UserManager, s *uman.Session) {
	if s.IsLogged() {
		c.HTML(http.StatusOK, "dashboard.tpl", gin.H{
			"username": s.User,
		})
	} else {
		c.HTML(http.StatusOK, "index.tpl", nil)
	}
}

func logout(c *gin.Context, um *uman.UserManager, s *uman.Session) {
	s.Logout()
	http.Redirect(c.Writer, c.Request, "/", 302)
}

func login(c *gin.Context, um *uman.UserManager, s *uman.Session) {
	if !s.IsLogged() {
		if !um.Login(c.PostForm("user"), c.PostForm("pass"), s) {
			c.HTML(http.StatusOK, "index.tpl", gin.H{
				"credError": true,
			})

			return
		}
	}

	http.Redirect(c.Writer, c.Request, "/", 302)
}

func (w *Web) routes() {
	w.engine.GET("/", func(c *gin.Context) {
		home(c, w.um, w.um.GetHTTPSession(c.Writer, c.Request))
	})

	w.engine.GET("/logout", func(c *gin.Context) {
		logout(c, w.um, w.um.GetHTTPSession(c.Writer, c.Request))
	})

	w.engine.POST("/login", func(c *gin.Context) {
		login(c, w.um, w.um.GetHTTPSession(c.Writer, c.Request))
	})
}