package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kgin"
)

//go:embed www
var www embed.FS

//go:embed test.html
var test string

func main() {
	app := kgin.Default()

	app.Use(kgin.DirHandler("/", "_example/web/www"))
	app.Use(kgin.StaticHandler("/www", www, "www"))
	kgin.GroupHandler(app.Group("/demo"), www, "www")
	kgin.GroupHandler(app, www, "www")

	app.GET("/test", kgin.HTMLHandler(test))
	app.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "NO FOUND...")
	})
	kgin.Root(app, www, "www")
	app.Run("0.0.0.0:8080")
}
