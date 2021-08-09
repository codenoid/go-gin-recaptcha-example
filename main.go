package main

import (
	"html/template"
	"net/http"
	"time"

	ginrecaptcha "github.com/codenoid/gin-recaptcha"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

var recaptchaSecret = "6LenMekbAAAAAIUbHoSiOmf1CkhECk75AcKUysRF"

func main() {
	r := gin.Default()
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/master",
		Funcs:        template.FuncMap{},
		DisableCache: true,
	})

	secret := "6LenMekbAAAAAIUbHoSiOmf1CkhECk75AcKUysRF"
	captcha, err := ginrecaptcha.InitRecaptchaV3(secret, 10*time.Second)
	if err != nil {
		panic(err)
	}

	captcha.SetErrResponse(func(c *gin.Context) {
		c.String(http.StatusUnprocessableEntity, "captcha error")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{})
	})

	r.POST("/", captcha.UseCaptcha, func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{"name": c.PostForm("name")})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
