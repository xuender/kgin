package kgin

import (
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Root(eng *gin.Engine, fsys fs.FS, dirs ...string) {
	eng.Use(func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodGet {
			if url := ctx.Request.URL.String(); url != "" && url != "/" {
				ctx.Header(CacheControl, MaxAge1y)
			}
		}
	})

	if len(dirs) > 0 {
		fsys, _ = fs.Sub(fsys, filepath.Join(dirs...))
	}

	eng.StaticFS("", http.FS(fsys))
}
