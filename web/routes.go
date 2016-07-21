package web

import (
	"github.com/gin-gonic/gin"
	"github.com/niemal/uman"
	"net/http"
	"os"
)

func css(c *gin.Context) {
	path := "web/css/" + c.Param("filename")

	if _, err := os.Stat(path); err == nil {
		c.Status(200)
		c.Header("Content-Type", "text/css")
		c.File(path)
	} else {
		c.String(404, "404 Not Found")
	}
}

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
	c.Redirect(302, "/")
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

	c.Redirect(302, "/")
}

func (w *Web) routes() {
	w.engine.GET("/css/:filename", func(c *gin.Context) {
		css(c)
	})

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