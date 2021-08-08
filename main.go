package main

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
)

var _ = os.Setenv("RECAPTCHA_SECRET", "6LenMekbAAAAAIUbHoSiOmf1CkhECk75AcKUysRF")
var recaptchaSecret = os.Getenv("RECAPTCHA_SECRET")

func main() {
	r := gin.Default()
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Funcs:        template.FuncMap{},
		DisableCache: true,
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{})
	})

	r.POST("/", func(c *gin.Context) {

		captcha, err := recaptcha.NewReCAPTCHA(recaptchaSecret, recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)
		if err != nil {
			c.HTML(http.StatusUnprocessableEntity, "home", gin.H{"error": err.Error()})
			return
		}

		if err := captcha.Verify(c.PostForm("g-recaptcha-response")); err != nil {
			c.HTML(http.StatusUnprocessableEntity, "home", gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusUnprocessableEntity, "home", gin.H{"name": c.PostForm("name")})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
