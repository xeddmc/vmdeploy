package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (w *Web) Home() {
	w.engine.GET("/", func(c *gin.Context) {
		session := w.um.GetHTTPSession(c.Writer, c.Request)
		logged := session.IsLogged()

		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"logged": logged,
		})
	})
}

func (w *Web) LogHandler() {
	w.engine.GET("/logout", func(c *gin.Context) {
		w.um.GetHTTPSession(c.Writer, c.Request).Logout()
		http.Redirect(c.Writer, c.Request, "/", 302)
	})

	w.engine.POST("/login", func(c *gin.Context) {
		if s := w.um.GetHTTPSession(c.Writer, c.Request); !s.IsLogged() {
			w.um.Login(c.PostForm("user"), c.PostForm("pass"), s)
		}

		http.Redirect(c.Writer, c.Request, "/", 302)
	})
}