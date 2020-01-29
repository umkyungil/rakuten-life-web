package main

import (
	"github.com/gin-gonic/gin"
	router "rakuten-life-web/server/adapter"
)

func main() {
	//mysql.InitGAE()
	//mysql.InitDev()
	r := router.Router(true)
	//http.Handle("/", r)
	//r.Run("8080")

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	r.Run()
}
