package main

import (
	"embed"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kgin"
	"github.com/xuender/kit/los"
)

//go:embed www
var www embed.FS

//go:embed test.html
var test string

func main() {
	app := gin.Default()
	app.GET("/test", kgin.HTMLHandler(test))
	_ = kgin.Embed(&app.RouterGroup, www, "www")
	los.Must0(app.Run("0.0.0.0:8080"))
}
