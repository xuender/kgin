package main

import (
	"embed"

	"github.com/xuender/kgin"
)

//go:embed www
var www embed.FS

//go:embed test.html
var test string

func main() {
	app := kgin.Default()

	app.Use(kgin.DirHandler("/", "_example/web/www"))
	app.Use(kgin.StaticHandler("/w", www, "www"))

	app.GET("/test", kgin.HTMLHandler(test))
	app.Run("0.0.0.0:8080")
}
