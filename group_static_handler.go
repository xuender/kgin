package kgin

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GroupHandler fs.
func GroupHandler(group *gin.RouterGroup, fsys fs.FS, dirs ...string) {
	var (
		base = group.BasePath()
		bas1 = base + "/"
	)

	group.Use(func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodGet {
			if url := ctx.Request.URL.String(); url != bas1 && url != base {
				ctx.Header(CacheControl, MaxAge1y)
			}
		}
	})

	group.StaticFS("", FileSystem(fsys, dirs...))
}
